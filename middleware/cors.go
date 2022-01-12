package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Cors() gin.HandlerFunc {
	return func(_ *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			//AllowOriginFunc:        nil,
			//AllowOriginRequestFunc: nil,
			AllowedMethods: []string{"*"},
			AllowedHeaders: []string{"Origin"},
			ExposedHeaders: []string{"Content-Length", "Authorization"},
			MaxAge:         12 * 60 * 60,
			//AllowCredentials:     false,
			//OptionsPassthrough:   false,
			//OptionsSuccessStatus: 0,
			//Debug:                false,
		})
	}
}
