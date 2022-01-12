package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)
	if err != nil {
		fmt.Println("写入日志失败！")
	}
	logger := logrus.New()
	//同时输出到日志文件里以及终端，这种方式保存进文件时有点问题
	logger.SetOutput(io.MultiWriter(os.Stdout, src))
	//只输出到文件里
	//logger.Out = src
	logger.SetLevel(logrus.TraceLevel)
	//设置为json格式
	logger.SetFormatter(&logrus.JSONFormatter{})
	//logger.SetReportCaller(true)
	//按时间分割日志
	linkName := "latest_log.log"
	logWriter, _ := rotatelog.New(
		filePath+"%Y%m%d.log",
		rotatelog.WithMaxAge(7*24*time.Hour),
		rotatelog.WithRotationTime(24*time.Hour), //分割时间
		rotatelog.WithLinkName(linkName),
	)
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.PanicLevel: logWriter,
		logrus.WarnLevel:  logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		SpendTime := fmt.Sprintf("%v ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
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
		reqPath := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  HostName,
			"status":    statusCode,
			"SpendTime": SpendTime,
			"IP":        clientIP,
			"Method":    reqMethod,
			"Path":      reqPath,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
