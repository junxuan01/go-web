package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv 尝试加载项目根目录下的 .env 文件（不存在则忽略）
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		// 本地没有 .env 时无需报错，通常在容器/CI用环境变量注入
		log.Println("未发现 .env 文件，跳过加载（可用环境变量配置数据库）")
	} else {
		log.Println("已加载 .env 环境变量")
	}
}
