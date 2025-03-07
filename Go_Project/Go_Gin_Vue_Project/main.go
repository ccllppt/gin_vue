package main

import (
	"Go_Gin_Vue_Project/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	r := gin.Default()
	r = CollectRoutes(r)
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
	r.Run("0.0.0.0:8080")
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
