package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// [
//
//	"id",
//	"name",
//	"trackName",
//	"artistName",
//	"albumName",
//	"duration",
//	"instrumental",
//	"plainLyrics",
//	"syncedLyrics",
//	"lyricsfile"
//
// ]
// LrclibSearchResponse 定义 LRCLIB 搜索接口的返回结构
type LrclibSearchResponse []struct {
	ID           int     `json:"id"`
	TrackName    string  `json:"trackName"`
	ArtistName   string  `json:"artistName"`
	AlbumName    string  `json:"albumName"`
	Duration     float64 `json:"duration"`
	SyncedLyrics string  `json:"syncedLyrics"` // 带时间轴的 LRC 歌词
	PlainLyrics  string  `json:"plainLyrics"`  // 纯文本歌词
}

// FetchLrcFromNetwork 从 LRCLIB 获取歌词
func FetchLrcFromNetwork(title string, artist string) (string, error) {
	time.Sleep(500 * time.Millisecond)
	if title == "" {
		return "", fmt.Errorf("缺少歌曲名称，无法搜索")
	}

	// 1. 拼接搜索 API（使用 search 接口容错率更高）
	baseURL := "https://lrclib.net/api/search"
	queryURL := fmt.Sprintf("%s?track_name=%s&artist_name=%s", baseURL, url.QueryEscape(title), url.QueryEscape(artist))
	// 2. 发起 HTTP 请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		return "", err
	}
	// LRCLIB 官方建议携带一个 User-Agent 标明身份
	req.Header.Set("User-Agent", "fnMusicServe/1.0 (NAS Music Player)")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("网络请求失败，状态码: %d", resp.StatusCode)
	}

	// 3. 解析 JSON 数据
	body, _ := io.ReadAll(resp.Body)
	var searchResults LrclibSearchResponse
	if err := json.Unmarshal(body, &searchResults); err != nil {
		return "", err
	}

	// 4. 提取歌词（优先获取带时间轴的 SyncedLyrics）
	if len(searchResults) > 0 {
		bestMatch := searchResults[0]
		if bestMatch.SyncedLyrics != "" {
			return bestMatch.SyncedLyrics, nil
		}
		if bestMatch.PlainLyrics != "" {
			return bestMatch.PlainLyrics, nil
		}
	}

	return "", fmt.Errorf("未找到匹配的歌词")
}
