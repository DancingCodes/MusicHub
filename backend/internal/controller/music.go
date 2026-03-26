package controller

import (
	"backend/internal/service"
	"backend/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNetMusicList(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		utils.Error(c, "请输入搜索关键词")
		return
	}

	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	list, total, err := service.SearchNetease(name, pageNo, pageSize)
	if err != nil {
		utils.Error(c, "搜索失败: "+err.Error())
		return
	}

	utils.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func SaveMusic(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数校验失败，请传入有效的 id")
		return
	}

	music, err := service.SaveMusicLogic(req.ID)
	if err != nil {
		utils.Error(c, err.Error())
		return
	}

	utils.Success(c, music)
}

func GetMusicList(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	list, total, err := service.GetMusicListLogic(pageNo, pageSize)
	if err != nil {
		utils.Error(c, "获取列表失败: "+err.Error())
		return
	}

	utils.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}
