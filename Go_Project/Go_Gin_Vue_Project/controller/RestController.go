package controller

import "github.com/gin-gonic/gin"

// RestController 是 RESTful 控制器的接口，定义了增删改查方法
type RestController interface {
	Create(c *gin.Context) // 创建资源
	Update(c *gin.Context) // 更新资源
	Show(c *gin.Context)   // 查询资源
	Delete(c *gin.Context) // 删除资源
}
