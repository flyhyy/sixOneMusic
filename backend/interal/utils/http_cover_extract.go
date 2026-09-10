package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ItunesSearchResponse 定义解析 iTunes JSON 的结构体
type ItunesSearchResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		ArtworkUrl100 string `json:"artworkUrl100"` // 100x100 尺寸的封面图
	} `json:"results"`
}

// FetchCoverFromNetwork 尝试从网络获取封面 URL
func FetchCoverFromNetwork(title string, artist string) (string, error) {
	// 1. 拼接搜索关键词 (例如: "晴天 周杰伦")
	query := url.QueryEscape(fmt.Sprintf("%s %s", title, artist))
	apiURL := fmt.Sprintf("https://itunes.apple.com/search?term=%s&entity=song&limit=1", query)

	// 2. 发起 HTTP GET 请求
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 3. 读取并解析 JSON 响应
	body, _ := io.ReadAll(resp.Body)
	var searchResult ItunesSearchResponse
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return "", err
	}

	// 4. 提取封面并替换为高清大图
	if searchResult.ResultCount > 0 && searchResult.Results[0].ArtworkUrl100 != "" {
		highResUrl := searchResult.Results[0].ArtworkUrl100
		// iTunes 默认返回 100x100，可以通过修改 URL 字符串获取 600x600 的高清大图
		highResUrl = strings.Replace(highResUrl, "100x100bb", "600x600bb", 1)
		return highResUrl, nil
	}

	return "", fmt.Errorf("未在网络上找到匹配的封面")
}
