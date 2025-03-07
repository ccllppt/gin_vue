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

func Register(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	var requestUser = Model.User{}
	c.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandmString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在

	if isTelePhoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//如果用户不存在就创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	newUser := Model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	DB.Create(&newUser)

	//发放token
	token, err := common.RelaeseToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "系统异常",
		})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "Register Success!")
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	var requestUser = Model.User{}
	c.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//判断手机号是否存在
	var user Model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "用户不存在",
		})
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "密码错误",
		})
		return
	}
	//发放token
	token, err := common.RelaeseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "系统异常",
		})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "login success")
}

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
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{"user": dto.ToUserDTOs(user.(Model.User))},
	})
}
