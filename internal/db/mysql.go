package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // 全局数据库连接（简单示例用法）

// InitMySQL 初始化MySQL连接，例如：
// DSN: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func InitMySQL() error {
	// 从环境变量读取配置，更贴近真实项目
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASS", "password")
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	name := getenv("DB_NAME", "go_web")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	glogger := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: glogger})
	if err != nil {
		return fmt.Errorf("连接MySQL失败: %w", err)
	}
	DB = db
	log.Println("MySQL 连接成功")
	return nil
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
