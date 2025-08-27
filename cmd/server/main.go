package main

import (
	"go-web/internal/config"
	"go-web/internal/routes"
	"log"
)

func main() {
	// 加载 .env（若存在）
	config.LoadEnv()
	// 设置路由
	router := routes.SetupRoutes()
	// 启动服务器，监听8080端口
	log.Println("服务器启动中，监听端口 :8080")
	log.Println("访问 http://localhost:8080/api/ping 测试服务")

	// 启动HTTP服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
