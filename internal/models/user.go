package models

import "time"

// User 用户结构体，用于表示用户信息
// 使用GORM标签与JSON标签，方便与数据库与API交互
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`                  // 用户ID（自增主键）
	Name      string    `json:"name" gorm:"type:varchar(64);not null"`               // 用户名
	Email     string    `json:"email" gorm:"type:varchar(128);uniqueIndex;not null"` // 邮箱（唯一索引）
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Response 通用响应结构体，用于统一API响应格式
type Response struct {
	Code    int    `json:"code"`    // 状态码，200表示成功
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 响应数据，可以是任意类型
}

// CreateUserRequest 创建用户请求结构体，用于接收POST请求参数
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`  // 用户名，必填
	Email string `json:"email" binding:"required"` // 邮箱，必填
}
