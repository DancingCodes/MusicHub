package router

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	music := r.Group("/music")
	{
		music.GET("/getNetMusicList", controller.GetNetMusicList)
		music.POST("/saveMusic", controller.SaveMusic)
	}

	return r
}
