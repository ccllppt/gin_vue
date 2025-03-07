package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//{
//	code:20001,
//	data:xxx,
//	msg:xx
//}

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
func Success(c *gin.Context, date gin.H, msg string) {
	Response(c, http.StatusOK, 200, date, msg)
}
func Fail(c *gin.Context, date gin.H, msg string) {
	Response(c, http.StatusOK, 400, date, msg)
}
