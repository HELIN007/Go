package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelog "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"math"
	"os"
	"time"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func init() {
	//filePath := "log/zaplog.log"
	//src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)
	//if err != nil {
	//	fmt.Println("写入zap日志失败！")
	//}
	encoderConfig := zap.NewProductionEncoderConfig()
	//使用大写彩色字母记录日志级别(json格式时彩色不生效)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	fileWriteSyncer := zapcore.AddSync(getWriter())
	//只写进log文件中
	//core := zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel)
	//写进log文件中并输出至终端
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	logger = zap.New(core)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("关闭zaplog错误，原因是：", err)
		}
	}(logger)
	//logger = zap.New(core, zap.AddCaller())
	//logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		HostName, err := os.Hostname()
		if err != nil {
			HostName = "未知用户"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		reqMethod := c.Request.Method
		//reqPath := c.Request.RequestURI
		c.Next()
		spendTime := time.Since(startTime)
		logger.Info(path,
			zap.String("hostname", HostName),
			zap.Int("status", statusCode),
			zap.String("clientIP", clientIP),
			zap.String("method", reqMethod),
			zap.String("agent", userAgent),
			//zap.String("url", reqPath),
			zap.String("cost", fmt.Sprintf("%v ms", int(math.Ceil(float64(spendTime.Nanoseconds()/1000000.0))))),
		)
	}
}

func getWriter() io.Writer {
	filePath := "log/log"
	linkName := "zap_latest_log.log"
	hook, err := rotatelog.New(
		filePath+"%Y%m%d.log",
		rotatelog.WithMaxAge(7*24*time.Hour),
		rotatelog.WithRotationTime(24*time.Hour),
		rotatelog.WithLinkName(linkName),
	)
	if err != nil {
		panic(nil)
	}
	return hook
}

// Infof 格式化输出info日志
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}
