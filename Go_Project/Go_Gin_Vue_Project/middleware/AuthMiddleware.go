package middleware

import (
	"Go_Gin_Vue_Project/Model"
	"Go_Gin_Vue_Project/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 是一个 Gin 中间件，用于验证用户权限
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Token
		tokenString := c.GetHeader("Authorization")
		// 检查 Token 是否存在且格式正确
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不够",
			})
			c.Abort() // 终止请求
			return
		}
		// 去掉 "Bearer " 前缀，获取实际的 Token
		tokenString = tokenString[7:]

		// 解析并验证 Token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		// 从 Token 中获取用户 ID
		userId := claims.UserId
		DB := common.GetDB() // 获取数据库连接
		var user Model.User
		// 根据用户 ID 查询用户信息
		DB.First(&user, userId)

		// 如果用户不存在，返回权限不足
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中，供后续处理使用
		c.Set("user", user)
		c.Next() // 继续处理请求
	}
}
