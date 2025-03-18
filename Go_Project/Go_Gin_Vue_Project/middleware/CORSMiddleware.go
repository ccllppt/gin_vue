package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware 是一个 Gin 中间件，用于处理跨域请求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的请求来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许携带凭证
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// 如果是 OPTIONS 请求，直接返回 204 状态码
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next() // 继续处理请求
	}
}
