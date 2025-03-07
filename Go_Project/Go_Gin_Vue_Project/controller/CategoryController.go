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

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(Model.Category{})

	return CategoryController{DB: db}
}

func (ctx CategoryController) Create(c *gin.Context) {
	var requestCategory Model.Category
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "请求参数错误")
		return
	}
	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}
	if err := ctx.DB.Create(&requestCategory).Error; err != nil {
		response.Fail(c, nil, "创建分类失败")
		return
	}
	response.Success(c, gin.H{"category": requestCategory}, "分类创建成功")
}
func (ctx CategoryController) Show(c *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	var category Model.Category
	if err := ctx.DB.First(&category, categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "分类不存在")
		}
		response.Fail(c, nil, "查询分类失败")
		return
	}
	response.Success(c, gin.H{"category": category}, "成功查看")
}

func (ctx CategoryController) Update(c *gin.Context) {
	//绑定body中的参数
	var requestCategory Model.Category
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "请求参数错误")
		return
	}
	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}
	//获取path中的参数
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	var updateCategory Model.Category
	if err := ctx.DB.First(&updateCategory, categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "分类不存在")
		}
		response.Fail(c, nil, "查询分类失败")
		return
	}
	//更新分类
	ctx.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(c, gin.H{"category": updateCategory}, "修改成功")
}
func (ctx CategoryController) Delete(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c, nil, "无效的分类ID")
		return
	}

	var category Model.Category
	if err := ctx.DB.First(&category, categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "分类不存在")
			return
		}
		response.Fail(c, nil, "查询分类失败")
		return
	}

	if err := ctx.DB.Delete(&category).Error; err != nil {
		response.Fail(c, nil, "删除失败，请重试")
		return
	}
	response.Success(c, nil, "删除成功")
}
