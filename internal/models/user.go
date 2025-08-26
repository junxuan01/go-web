package models

// User 用户结构体，用于表示用户信息
type User struct {
	ID    int    `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
	Email string `json:"email"` // 用户邮箱
}

// Response 通用响应结构体，用于统一API响应格式
type Response struct {
	Code    int         `json:"code"`    // 状态码，200表示成功
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据，可以是任意类型
}

// CreateUserRequest 创建用户请求结构体，用于接收POST请求参数
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`  // 用户名，必填
	Email string `json:"email" binding:"required"` // 邮箱，必填
}
