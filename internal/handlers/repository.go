package handlers

import (
	"errors"
	"go-web/internal/db"
	"go-web/internal/models"
)

// UserRepository 定义用户数据访问接口，便于替换实现与单元测试
type UserRepository interface {
	List() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(u *models.User) error
}

// GormUserRepository 基于GORM的实现
type GormUserRepository struct{}

func NewGormUserRepository() *GormUserRepository { return &GormUserRepository{} }

func (r *GormUserRepository) List() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Create(u *models.User) error {
	if u == nil {
		return errors.New("nil user")
	}
	return db.DB.Create(u).Error
}
