package router

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getNetMusicList", handler.GetNetMusicList)
	r.POST("/saveMusic", handler.SaveMusic)
	return r
}
