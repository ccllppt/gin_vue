package main

import (
	"Go_Gin_Vue_Project/controller"
	"Go_Gin_Vue_Project/middleware"
	"github.com/gin-gonic/gin"
)

// CollectRoutes 注册所有路由并返回 Gin 引擎
func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 使用 CORS 中间件，允许跨域请求
	r.Use(middleware.CORSMiddleware())

	// 用户认证相关路由
	r.POST("/api/auth/register", controller.Register)                     // 用户注册
	r.POST("/api/auth/login", controller.Login)                           // 用户登录
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) // 获取当前用户信息（需要认证）

	// 分类相关路由
	categoryRoutes := r.Group("/category")                   // 创建分类路由组
	categoryController := controller.NewCategoryController() // 创建分类控制器实例

	// 分类路由组的具体路由
	categoryRoutes.POST("", categoryController.Create)       // 创建分类
	categoryRoutes.PUT("/:id", categoryController.Update)    // 更新分类
	categoryRoutes.GET("/:id", categoryController.Show)      // 查询分类
	categoryRoutes.DELETE("/:id", categoryController.Delete) // 删除分类

	return r // 返回注册了路由的 Gin 引擎
}
