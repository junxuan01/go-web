package handlers

import (
	"net/http"
	"strconv"

	"go-web/internal/models"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器结构体，包含所有用户相关的处理方法
type UserHandler struct {
	repo UserRepository // 数据访问仓储
}

// NewUserHandler 创建并返回一个新的用户处理器实例
func NewUserHandler(repo UserRepository) *UserHandler { return &UserHandler{repo: repo} }

// GetUsers 处理GET /users请求，返回所有用户列表
func (h *UserHandler) GetUsers(c *gin.Context) {
	// 创建统一响应格式
	users, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "查询用户失败: " + err.Error()})
		return
	}
	response := models.Response{Code: 200, Message: "获取用户列表成功", Data: users}

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
	user, err := h.repo.GetByID(uint(id))
	if err == nil && user != nil {
		c.JSON(http.StatusOK, models.Response{Code: 200, Message: "获取用户成功", Data: user})
		return
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
	newUser := models.User{Name: req.Name, Email: req.Email}
	if err := h.repo.Create(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "创建用户失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.Response{Code: 201, Message: "用户创建成功", Data: newUser})
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
