package controller

import (
	"Go_Gin_Vue_Project/Model"
	"Go_Gin_Vue_Project/common"
	"Go_Gin_Vue_Project/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// ICategoryController 是 CategoryController 的接口，定义了 RESTful 方法
type ICategoryController interface {
	RestController
}

// CategoryController 是分类控制器，实现了 ICategoryController 接口
type CategoryController struct {
	DB *gorm.DB // 数据库连接实例
}

// NewCategoryController 创建并返回一个新的 CategoryController 实例
func NewCategoryController() ICategoryController {
	db := common.GetDB()             // 获取全局数据库连接
	db.AutoMigrate(Model.Category{}) // 自动迁移 Category 表结构

	return CategoryController{DB: db}
}

// Create 创建分类
func (ctx CategoryController) Create(c *gin.Context) {
	var requestCategory Model.Category
	// 绑定请求体中的参数到 requestCategory
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "请求参数错误")
		return
	}
	// 验证分类名称是否为空
	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}
	// 在数据库中创建分类
	if err := ctx.DB.Create(&requestCategory).Error; err != nil {
		response.Fail(c, nil, "创建分类失败")
		return
	}
	// 返回成功响应
	response.Success(c, gin.H{"category": requestCategory}, "分类创建成功")
}

// Show 根据 ID 查询分类
func (ctx CategoryController) Show(c *gin.Context) {
	// 从 URL 路径中获取分类 ID
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	var category Model.Category
	// 根据 ID 查询分类
	if err := ctx.DB.First(&category, categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "分类不存在")
		}
		response.Fail(c, nil, "查询分类失败")
		return
	}
	// 返回成功响应
	response.Success(c, gin.H{"category": category}, "成功查看")
}

// Update 更新分类
func (ctx CategoryController) Update(c *gin.Context) {
	// 绑定请求体中的参数到 requestCategory
	var requestCategory Model.Category
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "请求参数错误")
		return
	}
	// 验证分类名称是否为空
	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}
	// 从 URL 路径中获取分类 ID
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	var updateCategory Model.Category
	// 根据 ID 查询分类
	if err := ctx.DB.First(&updateCategory, categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "分类不存在")
		}
		response.Fail(c, nil, "查询分类失败")
		return
	}
	// 更新分类名称
	ctx.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	// 返回成功响应
	response.Success(c, gin.H{"category": updateCategory}, "修改成功")
}

// Delete 删除分类
func (ctx CategoryController) Delete(c *gin.Context) {
	// 从 URL 路径中获取分类 ID
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c, nil, "无效的分类ID")
		return
	}

	var category Model.Category
	// 根据 ID 查询分类
	if err := ctx.DB.First(&category, categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "分类不存在")
			return
		}
		response.Fail(c, nil, "查询分类失败")
		return
	}

	// 删除分类
	if err := ctx.DB.Delete(&category).Error; err != nil {
		response.Fail(c, nil, "删除失败，请重试")
		return
	}
	// 返回成功响应
	response.Success(c, nil, "删除成功")
}
