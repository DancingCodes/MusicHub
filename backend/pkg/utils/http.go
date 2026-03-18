package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

// GetJSON 支持传入 headers，不需要传则传 nil
func GetJSON[T any](url string, headers map[string]string) (T, error) {
	var result T

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	// 设置 Header
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode >= 400 {
		return result, fmt.Errorf("http error: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func DownloadToFile(url string, headers map[string]string, dest string) error {
	// 1. 确保本地目录存在
	dir := filepath.Dir(dest)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("无法创建目录 %s: %v", dir, err)
	}

	// 2. 创建请求
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// 注入 Headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 3. 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，响应码: %d", resp.StatusCode)
	}

	// 4. 创建本地文件
	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("无法创建文件 %s: %v", dest, err)
	}
	defer out.Close()

	// 5. 流式写入，避免把大文件全部加载到内存
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	return nil
}
