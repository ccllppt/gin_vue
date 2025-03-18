package common

import (
	"Go_Gin_Vue_Project/Model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

// DB 是一个全局变量，用于保存数据库连接实例
var DB *gorm.DB

// InitDB 初始化数据库连接，并返回一个 *gorm.DB 实例
func InitDB() *gorm.DB {
	// 从配置文件中读取数据库连接信息
	host := viper.GetString("datasource.host")         // 数据库主机地址
	database := viper.GetString("datasource.database") // 数据库名称
	username := viper.GetString("datasource.username") // 数据库用户名
	password := viper.GetString("datasource.password") // 数据库密码
	charset := viper.GetString("datasource.charset")   // 数据库字符集
	loc := viper.GetString("datasource.loc")           // 数据库时区

	// 拼接数据库连接字符串
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		database,
		charset,
		url.QueryEscape(loc))

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error()) // 如果连接失败，直接 panic
	}

	// 自动迁移 User 表结构
	if err := db.AutoMigrate(&Model.User{}); err != nil {
		panic("failed to auto migrate database" + err.Error()) // 如果迁移失败，直接 panic
	}

	fmt.Println("Database connected and table ready") // 打印连接成功信息
	DB = db                                           // 将数据库连接实例赋值给全局变量 DB
	return db
}

// GetDB 返回全局的数据库连接实例
func GetDB() *gorm.DB {
	return DB
}
