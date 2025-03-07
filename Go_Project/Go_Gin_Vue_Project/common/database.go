package common

import (
	"Go_Gin_Vue_Project/Model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	//拼接连接字符串
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	if err := db.AutoMigrate(&Model.User{}); err != nil {
		panic("failed to auto migrate database" + err.Error())
	}
	fmt.Println("Database connected and table ready")
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
