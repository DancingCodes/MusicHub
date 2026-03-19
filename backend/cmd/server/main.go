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

	saveMusicPath := os.Getenv("SAVE_MUSIC_DIR")
	err := os.MkdirAll(saveMusicPath, 0755)
	if err != nil {
		log.Fatalf("无法创建上传目录 [%s]: %v", saveMusicPath, err)
	}

	db.InitDB()
	r := router.SetupRouter()

	port := os.Getenv("PORT")

	if err := r.Run(port); err != nil {
		log.Fatalf("❌ 服务启动崩溃: %v", err)
	}
}
