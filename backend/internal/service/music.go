package service

import (
	"backend/internal/db"
	"backend/internal/external/dto"
	"backend/internal/model"
	"backend/pkg/utils"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

func SearchNetease(name string, pageNo, pageSize int) ([]model.Music, int64, error) {
	params := url.Values{}
	params.Set("s", name)
	params.Set("type", "1")
	params.Set("offset", strconv.Itoa((pageNo-1)*pageSize))
	params.Set("limit", strconv.Itoa(pageSize))

	apiUrl := "https://music.163.com/api/search/get/web?" + params.Encode()

	type rawSong struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
	}
	type NeteaseRawResponse struct {
		Result struct {
			Songs     []rawSong `json:"songs"`
			SongCount int64     `json:"songCount"`
		} `json:"result"`
	}
	raw, err := utils.GetJSON[NeteaseRawResponse](apiUrl, nil)
	if err != nil {
		return nil, 0, err
	}
	var musicList []model.Music
	for _, rs := range raw.Result.Songs {
		var artistNames []string
		for _, a := range rs.Artists {
			artistNames = append(artistNames, a.Name)
		}

		musicList = append(musicList, model.Music{
			ID:      rs.ID,
			Name:    rs.Name,
			Artists: strings.Join(artistNames, ","),
		})
	}

	if musicList == nil {
		musicList = []model.Music{}
	}

	return musicList, raw.Result.SongCount, nil
}

func SaveMusicLogic(songID int) (*model.Music, error) {
	var existing model.Music
	if err := db.DB.Where("id = ?", songID).First(&existing).Error; err == nil {
		return &existing, nil
	}

	var (
		wg            sync.WaitGroup
		res1          dto.NETEASEMusicDetailsStruct
		res2          dto.NETEASEMusicLyricStruct
		res3          dto.NETEASEMusicURLStruct
		err1, _, err3 error
	)

	wg.Add(3)
	// 获取详情
	go func() {
		defer wg.Done()
		apiUrl := fmt.Sprintf("https://music.163.com/api/v3/song/detail?id=%d&c=[{id:%d}]", songID, songID)
		res1, err1 = utils.GetJSON[dto.NETEASEMusicDetailsStruct](apiUrl, nil)
	}()
	// 获取歌词
	go func() {
		defer wg.Done()
		apiUrl := fmt.Sprintf("https://music.163.com/api/song/lyric?id=%d&lv=-1&tv=-1", songID)
		res2, _ = utils.GetJSON[dto.NETEASEMusicLyricStruct](apiUrl, nil)
	}()
	// 获取播放地址
	go func() {
		defer wg.Done()
		apiUrl := fmt.Sprintf("https://music.163.com/api/song/enhance/player/url/v1?ids=[%d]&encodeType=aac&level=jymaster", songID)
		res3, err3 = utils.GetJSON[dto.NETEASEMusicURLStruct](apiUrl, map[string]string{
			"Cookie": os.Getenv("NETEASE_COOKIE"),
		})
	}()

	wg.Wait()

	if err1 != nil || len(res1.Songs) == 0 {
		return nil, fmt.Errorf("获取详情失败: %v", err1)
	}
	if err3 != nil || len(res3.Data) == 0 || res3.Data[0].URL == "" {
		return nil, fmt.Errorf("获取播放链接失败或版权受限")
	}

	song := res1.Songs[0]
	remoteUrl := res3.Data[0].URL
	fileType := res3.Data[0].Type

	localFileName := fmt.Sprintf("%d.%s", songID, fileType)
	localPath := os.Getenv("SAVE_MUSIC_DIR") + "/" + localFileName
	if err := utils.DownloadToFile(remoteUrl, localPath); err != nil {
		return nil, fmt.Errorf("下载文件失败: %v", err)
	}

	var artists []string
	for _, ar := range song.Ar {
		artists = append(artists, ar.Name)
	}

	// 5. 构造模型并入库
	newMusic := model.Music{
		ID:       uint(song.ID),
		Name:     song.Name,
		Url:      os.Getenv("MUSIC_BASE_URL") + localFileName,
		PicUrl:   song.Al.PicURL,
		Artists:  strings.Join(artists, ","),
		Duration: song.Dt,
		Lyric:    res2.Lrc.Lyric,
	}

	if err := db.DB.Create(&newMusic).Error; err != nil {
		return nil, fmt.Errorf("数据库入库失败: %v", err)
	}

	return &newMusic, nil
}

func GetMusicListLogic(pageNo, pageSize int) ([]model.Music, int64, error) {
	var musicList []model.Music
	var total int64

	if err := db.DB.Model(&model.Music{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNo - 1) * pageSize
	err := db.DB.Offset(offset).Limit(pageSize).Order("id desc").Find(&musicList).Error

	return musicList, total, err
}
