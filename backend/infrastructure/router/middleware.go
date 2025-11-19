package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS は開発環境で全てのリクエストを許可するCORSミドルウェアを返す
func CORS() gin.HandlerFunc {
	// CORS設定 (開発用: 全オリジン許可)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	config.AddAllowHeaders("Authorization", "Content-Type")

	return cors.New(config)
}
