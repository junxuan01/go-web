package routes

import (
	"go-web/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有的路由规则
func SetupRoutes() *gin.Engine {
	// 创建Gin路由引擎，使用默认中间件（logger和recovery）
	router := gin.Default()

	// 创建用户处理器实例
	userHandler := handlers.NewUserHandler()

	// 设置API路由组，所有API都以/api开头
	api := router.Group("/api")
	{
		// 健康检查接口
		api.GET("/ping", userHandler.Ping)

		// 用户相关接口
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetUsers)        // GET /api/users - 获取所有用户
			users.GET("/:id", userHandler.GetUserByID) // GET /api/users/:id - 根据ID获取用户
			users.POST("", userHandler.CreateUser)     // POST /api/users - 创建新用户
		}
	}

	return router
}
