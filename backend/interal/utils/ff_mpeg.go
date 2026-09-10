package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"fnMusicServe/interal/model"

	"github.com/dhowden/tag"
)

type FFprobeOutput struct {
	Format FFprobeFormat `json:"format"`
}

type FFprobeFormat struct {
	Filename string            `json:"fileName"`
	Duration string            `json:"duration"` // ffprobe 返回的时长是带小数点的字符串，例如 "234.567"
	Size     string            `json:"size"`     // 返回的也是字符串字节数
	Tags     map[string]string `json:"tags"`     // ID3 等元数据都在这个 map 里
}

// generateMD5 辅助函数：生成字符串的 MD5 哈希值
func generateMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ParseAudioWithFFprobe(filePath string, lrcExist bool) (*model.SongDB, error) {
	// 1. 构建执行命令
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json",
		"-show_format", "-show_streams", filePath)

	// 2. 获取命令执行的标准输出
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("ffprobe 执行失败: %v, 请检查是否安装了 ffmpeg 并且文件路径正确", err)
	}

	// 3. 解析 JSON 结果
	var probeData FFprobeOutput
	err = json.Unmarshal(out.Bytes(), &probeData)
	if err != nil {
		return nil, fmt.Errorf("解析 ffprobe json 失败: %v", err)
	}

	// 4. 数据类型转换
	var durationInt int
	if probeData.Format.Duration != "" {
		durationFloat, _ := strconv.ParseFloat(probeData.Format.Duration, 64)
		durationInt = int(durationFloat)
	}

	var sizeInt64 int64
	if probeData.Format.Size != "" {
		sizeInt64, _ = strconv.ParseInt(probeData.Format.Size, 10, 64)
	}

	// 5. 提取 Tags 中的元数据
	tags := probeData.Format.Tags
	title := getTagValue(tags, "title")
	singer := getTagValue(tags, "artist")
	album := getTagValue(tags, "album")
	style := getTagValue(tags, "genre")
	lrc := getTagValue(tags, "lyrics")

	// 如果没有读取到歌曲名，使用安全的 filepath.Base 获取文件名兜底
	if title == "" {
		fileNameWithExt := filepath.Base(filePath)
		ext := filepath.Ext(fileNameWithExt)
		title = strings.TrimSuffix(fileNameWithExt, ext)
	}

	lrcPath := ""
	if lrcExist {
		if lrc == "" {
			newLrc, err := FetchLrcFromNetwork(title, singer)
			if err == nil && newLrc != "" {
				lrc = newLrc
				log.Printf("成功从网络下载歌词: %s - %s", singer, title)
			}
		}
		if lrc != "" {
			// 目前保持你的原有逻辑：存放在歌曲同级目录
			ext := filepath.Ext(filePath)
			lrcPath = strings.TrimSuffix(filePath, ext) + ".lrc"

			err := os.WriteFile(lrcPath, []byte(lrc), 0o644)
			if err != nil {
				log.Printf("歌词路径写入失败：%v\n", err)
			}
		}
		log.Printf("========lrc=============%v", lrc)

	}

	// ================= 封面提取逻辑：使用 Go 原生库写入统一文件夹 =================
	coverDir := filepath.Join("music", "covers") // 目标文件夹：项目根目录/music/covers
	coverPath := ""

	// 确保 covers 文件夹存在
	if err := os.MkdirAll(coverDir, os.ModePerm); err != nil {
		log.Printf("创建封面目录失败: %v\n", err)
	} else {
		// 使用文件绝对路径的 MD5 作为封面名
		coverFileName := generateMD5(filePath) + ".jpg"
		targetCoverPath := filepath.Join(coverDir, coverFileName)

		// 检查封面是否已经存在
		if _, err := os.Stat(targetCoverPath); os.IsNotExist(err) {

			// 1. 打开音频文件
			file, err := os.Open(filePath)
			if err == nil {
				// 2. 解析音频标签
				m, err := tag.ReadFrom(file)
				// log.Printf("=========m===========%v", m)
				// log.Printf("=========merr===========%v", err)

				if err == nil && m != nil {
					// 3. 获取封面图片数据
					pic := m.Picture()
					log.Printf("=========pic===========%v", pic)

					// 确保图片数据存在且不为空
					if pic != nil && len(pic.Data) > 0 {
						// 4. 将纯二进制图片数据直接写入本地文件
						err = os.WriteFile(targetCoverPath, pic.Data, 0o644)
						if err != nil {
							log.Printf("封面写入失败: %v\n", err)
						} else {
							// 写入成功，记录路径
							coverPath = targetCoverPath
						}
					} else {

						url, err := FetchCoverFromNetwork(title, singer)
						if err == nil {
							coverPath = url
						}
					}
				}
				file.Close() // 记得关闭文件
			}
		} else {
			// 如果文件已经存在，直接复用路径
			coverPath = targetCoverPath
		}
	}
	// ================================================================

	// 6. 组装返回你的数据库模型
	song := &model.SongDB{
		Title:     title,
		Singer:    singer,
		StyleName: style,
		AlbumName: album,
		FilePath:  filePath,
		Duration:  durationInt, // 精确的时长
		Size:      sizeInt64,   // 精确的文件大小
		LrcPath:   lrcPath,
		CoverPath: coverPath, // 存入如 "data/covers/e10adc3949ba59abbe56e057f20f883e.jpg"
	}
	return song, nil
}

func getTagValue(tags map[string]string, targetKey string) string {
	targetKey = strings.ToLower(targetKey)
	for k, v := range tags {
		if strings.ToLower(k) == targetKey {
			return v
		}
	}
	return ""
}
