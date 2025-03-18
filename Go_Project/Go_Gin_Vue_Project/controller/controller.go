package controller

import (
	"Go_Gin_Vue_Project/Model"
	"Go_Gin_Vue_Project/common"
	"Go_Gin_Vue_Project/dto"
	"Go_Gin_Vue_Project/response"
	"Go_Gin_Vue_Project/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	DB := common.GetDB() // 获取全局数据库连接
	// 绑定请求体中的参数到 requestUser
	var requestUser = Model.User{}
	c.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	// 如果用户名为空，生成一个随机用户名
	if len(name) == 0 {
		name = util.RandmString(10)
	}
	log.Println(name, telephone, password)

	// 判断手机号是否已存在
	if isTelePhoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	// 对密码进行哈希加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	// 创建新用户
	newUser := Model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	DB.Create(&newUser)

	// 生成并发放 Token
	token, err := common.RelaeseToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "系统异常",
		})
		log.Printf("token generate error:%v", err)
		return
	}
	// 返回成功响应
	response.Success(c, gin.H{"token": token}, "Register Success!")
}

// Login 用户登录
func Login(c *gin.Context) {
	DB := common.GetDB() // 获取全局数据库连接
	// 绑定请求体中的参数到 requestUser
	var requestUser = Model.User{}
	c.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}

	// 根据手机号查询用户
	var user Model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "用户不存在",
		})
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "密码错误",
		})
		return
	}

	// 生成并发放 Token
	token, err := common.RelaeseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "系统异常",
		})
		log.Printf("token generate error:%v", err)
		return
	}
	// 返回成功响应
	response.Success(c, gin.H{"token": token}, "login success")
}

// isTelePhoneExist 判断手机号是否已存在
func isTelePhoneExist(db *gorm.DB, telephone string) bool {
	if db == nil {
		log.Println("Database connection is nil")
		return false
	}

	var user Model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// Info 获取当前用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user") // 从上下文中获取当前用户
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{"user": dto.ToUserDTOs(user.(Model.User))}, // 返回用户信息
	})
}
