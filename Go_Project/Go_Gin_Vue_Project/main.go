package main

import (
	"Go_Gin_Vue_Project/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// main 是程序的入口函数
func main() {
	InitConfig()          // 初始化配置文件
	db := common.InitDB() // 初始化数据库连接
	defer func() {
		sqlDB, _ := db.DB() // 获取底层的 SQL DB 对象
		sqlDB.Close()       // 在程序退出时关闭数据库连接
	}()

	r := gin.Default()   // 创建一个默认的 Gin 引擎
	r = CollectRoutes(r) // 注册路由

	// 从配置文件中读取服务器端口，如果未配置则使用默认端口 8080
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	if err := r.Run(":" + port); err != nil {
		panic(err) // 如果启动失败，直接 panic
	}
	r.Run("0.0.0.0:8080") // 监听所有网络接口的 8080 端口
}

// InitConfig 初始化配置文件
func InitConfig() {
	workDir, _ := os.Getwd()                 // 获取当前工作目录
	viper.SetConfigName("application")       // 设置配置文件名（不需要扩展名）
	viper.SetConfigType("yml")               // 设置配置文件类型为 YAML
	viper.AddConfigPath(workDir + "/config") // 设置配置文件路径
	err := viper.ReadInConfig()              // 读取配置文件
	if err != nil {
		panic(err) // 如果读取失败，直接 panic
	}
}
