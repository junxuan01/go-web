package handlers

import (
	"net/http"
	"strconv"

	"go-web/internal/models"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器结构体，包含所有用户相关的处理方法
type UserHandler struct {
	users  []models.User // 内存中存储用户数据（实际项目中会使用数据库）
	nextID int           // 下一个用户ID
}

// NewUserHandler 创建并返回一个新的用户处理器实例
func NewUserHandler() *UserHandler {
	return &UserHandler{
		users:  make([]models.User, 0), // 初始化空的用户切片
		nextID: 1,                      // 从ID 1开始
	}
}

// GetUsers 处理GET /users请求，返回所有用户列表
func (h *UserHandler) GetUsers(c *gin.Context) {
	// 创建统一响应格式
	response := models.Response{
		Code:    200,
		Message: "获取用户列表成功",
		Data:    h.users,
	}

	// 返回JSON响应
	c.JSON(http.StatusOK, response)
}

// GetUserByID 处理GET /users/:id请求，根据ID获取特定用户
func (h *UserHandler) GetUserByID(c *gin.Context) {
	// 从URL参数中获取用户ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// ID格式错误
		response := models.Response{
			Code:    400,
			Message: "用户ID格式错误",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// 查找用户
	for _, user := range h.users {
		if user.ID == id {
			response := models.Response{
				Code:    200,
				Message: "获取用户成功",
				Data:    user,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// 用户不存在
	response := models.Response{
		Code:    404,
		Message: "用户不存在",
		Data:    nil,
	}
	c.JSON(http.StatusNotFound, response)
}

// CreateUser 处理POST /users请求，创建新用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	// 绑定JSON数据到结构体，自动验证必填字段
	if err := c.ShouldBindJSON(&req); err != nil {
		response := models.Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// 创建新用户
	newUser := models.User{
		ID:    h.nextID,
		Name:  req.Name,
		Email: req.Email,
	}

	// 将用户添加到内存存储中
	h.users = append(h.users, newUser)
	h.nextID++ // 递增ID

	// 返回创建成功的响应
	response := models.Response{
		Code:    201,
		Message: "用户创建成功",
		Data:    newUser,
	}
	c.JSON(http.StatusCreated, response)
}

// Ping 处理GET /ping请求，用于健康检查
func (h *UserHandler) Ping(c *gin.Context) {
	response := models.Response{
		Code:    200,
		Message: "pong",
		Data:    gin.H{"timestamp": "now"},
	}
	c.JSON(http.StatusOK, response)
}
