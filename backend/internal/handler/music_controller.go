package handler

import (
	"backend/internal/external/dto"
	utils2 "backend/internal/utils"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNetMusicList(c *gin.Context) {
	name := c.Query("name")
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	params := url.Values{}
	params.Set("s", name)
	params.Set("type", "1")
	params.Set("offset", strconv.Itoa((pageNo-1)*pageSize))
	params.Set("limit", strconv.Itoa(pageSize))

	apiUrl := "https://music.163.com/api/search/get/web?" + params.Encode()

	res, err := utils2.GetJSON[any](apiUrl, nil)
	if err != nil {
		utils2.Error(c, "第三方接口调用失败: "+err.Error())
		return
	}

	utils2.Success(c, res)
}

func SaveMusic(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils2.Error(c, "参数校验失败，请传入有效的 id")
		return
	}

	//1.如果已入库就直接返回

	// 获取歌曲详情: res?.songs?.[0]
	params1 := url.Values{}
	params1.Set("id", strconv.Itoa(req.ID))
	params1.Set("c", fmt.Sprintf("[{id: %d}]", req.ID))
	baseUrl1 := "https://music.163.com/api/v3/song/detail"
	apiUrl1 := baseUrl1 + "?" + params1.Encode()

	res1, err1 := utils2.GetJSON[dto.NETEASEMusicDetailsStruct](apiUrl1, nil)
	if err1 != nil {
		utils2.Error(c, "第三方接口调用失败1: "+err1.Error())
		return
	}

	// 获取歌曲歌词: res?.lrc?.lyric || ''
	params2 := url.Values{}
	params2.Set("id", strconv.Itoa(req.ID))
	params2.Set("lv", "-1")
	params2.Set("tv", "-1")

	baseUrl2 := "https://music.163.com/api/song/lyric"
	apiUrl2 := baseUrl2 + "?" + params2.Encode()

	res2, err2 := utils2.GetJSON[dto.NETEASEMusicLyricStruct](apiUrl2, nil)
	if err2 != nil {
		utils2.Error(c, "第三方接口调用失败2: "+err2.Error())
		return
	}

	// 获取歌曲网络链接: res?.data?.[0]?.url
	params3 := url.Values{}
	params3.Set("ids", fmt.Sprintf("[%d]", req.ID))
	params3.Set("encodeType", "aac")
	params3.Set("level", "jymaster")

	baseUrl3 := "https://music.163.com/api/song/enhance/player/url/v1"
	apiUrl3 := baseUrl3 + "?" + params3.Encode()

	res3, err3 := utils2.GetJSON[dto.NETEASEMusicURLStruct](apiUrl3, map[string]string{
		"Cookie": os.Getenv("NETEASE_COOKIE"),
	})
	if err3 != nil {
		utils2.Error(c, "第三方接口调用失败3: "+err3.Error())
		return
	}

	song := res1.Songs[0]
	lyric := res2.Lrc.Lyric
	url1 := res3.Data[0].URL

	// 1.保存url1到指定目录

	// 2.入库

	utils2.Success(c, gin.H{
		"id":       song.ID,
		"name":     song.Name,
		"picUrl":   song.Al.PicURL,
		"artists":  song.Ar,
		"duration": song.Dt,
		"lyric":    lyric,
		"url":      url1,
	})
}
