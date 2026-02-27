package main

import (
	"backend/internal/db"
	"backend/internal/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，将使用系统环境变量")
	}

	err := os.MkdirAll("./saveMusics", 0755)
	if err != nil {
		log.Fatalf("无法创建保存音乐目录: %v", err)
	}

	db.InitDB()
	r := router.NewRouter()

	port := os.Getenv("SERVICE_PORT")

	if port == "" {
		port = ":8080"
	} else if port[0] != ':' {
		port = ":" + port
	}

	if err := r.Run(port); err != nil {
		log.Fatalf("❌ 服务启动崩溃: %v", err)
	}
}
