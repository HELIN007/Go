package middleware

import (
	"GinProject/utils"
	"GinProject/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//TODO Redis解决token过期问题

// SetToken 生成Token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			Issuer:    "lin",             //签发人
		},
	}
	requestClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	//SignedString 接收一个[]byte
	token, err := requestClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// ParseToken 解析token
func ParseToken(token string) (*MyClaims, int) {
	//可能会过期
	setToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	Infof("解析token，状态是%v", err)
	if err != nil {
		fmt.Println("解析token失败！原因为: ", err)
		return nil, errmsg.ErrorTokenRuntime
	}
	if key, code := setToken.Claims.(*MyClaims); code && setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

var code int

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求头
		Infof("正在处理token")
		tokenHeader := c.Request.Header.Get("authorization")
		if tokenHeader == "" {
			code = errmsg.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//最多返回n个子字符串，最后一个字符串为未分割的余数
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		//校验token的格式
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//解析token
		key, parsedCode := ParseToken(checkToken[1])
		if parsedCode != errmsg.SUCCESS {
			code = parsedCode
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
