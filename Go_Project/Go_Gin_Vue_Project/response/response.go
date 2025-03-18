package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 是统一的响应函数，用于返回 JSON 格式的响应
func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code, // 状态码
		"data": data, // 数据
		"msg":  msg,  // 消息
	})
}

// Success 返回成功的响应
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

// Fail 返回失败的响应
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
