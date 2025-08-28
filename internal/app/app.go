package app

import (
	"go-web/internal/config"
	"go-web/internal/db"
	"go-web/internal/models"
	"go-web/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// App 应用程序结构体
type App struct {
	Router *gin.Engine
}

// NewApp 创建新的应用实例
func NewApp() *App {
	return &App{}
}

// Initialize 初始化应用程序
func (a *App) Initialize() error {
	// 1. 加载配置
	config.LoadEnv()
	log.Println("配置加载完成")

	// 2. 初始化数据库
	if err := a.initDatabase(); err != nil {
		return err
	}

	// 3. 设置路由
	a.Router = routes.SetupRoutes()
	log.Println("路由配置完成")

	return nil
}

// initDatabase 初始化数据库连接和迁移
func (a *App) initDatabase() error {
	// 连接数据库
	if err := db.InitMySQL(); err != nil {
		log.Printf("数据库连接失败: %v", err)
		return err
	}
	log.Println("数据库连接成功")

	// 执行迁移
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Printf("数据库迁移失败: %v", err)
		return err
	}
	log.Println("数据库迁移完成")

	return nil
}

// Run 启动应用程序
func (a *App) Run(addr string) error {
	log.Printf("服务器启动中，监听端口 %s", addr)
	log.Printf("访问 http://localhost%s/api/ping 测试服务", addr)

	return a.Router.Run(addr)
}
