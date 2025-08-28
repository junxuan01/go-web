package main

import (
	"go-web/internal/app"
	"log"
)

func main() {
	// 创建应用实例
	application := app.NewApp()

	// 初始化应用（配置、数据库、路由）
	if err := application.Initialize(); err != nil {
		log.Fatalf("应用初始化失败: %v", err)
	}

	// 启动服务器
	if err := application.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
