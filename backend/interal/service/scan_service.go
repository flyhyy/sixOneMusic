package service

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
	"fnMusicServe/interal/utils"
)

type ScanSerImpl struct {
	ScanRepo   repository.ScanRepoFace
	IsScanning bool         // 标记是否正在扫描
	mu         sync.RWMutex // 读写锁，保证并发安全
}

type ScanSerFace interface {
	Scan([]string)
	GetScanStatus() bool
}

func NewScanSer(repo repository.ScanRepoFace) ScanSerFace {
	return &ScanSerImpl{ScanRepo: repo}
}

func ScanDir(rootDir string, path string, result chan<- model.SongDB, wg *sync.WaitGroup) {
	defer wg.Done()
	// 1. 使用绝对路径读取目录，更安全
	fullDirPath := filepath.Join(rootDir, path)
	arr, err := os.ReadDir(fullDirPath)
	if err != nil {
		log.Printf("读取目录失败 %s 失败,%v\n", err, path)
		return
	}
	for _, item := range arr {
		if !item.IsDir() {
			// 2. 核心逻辑：只匹配并处理音频文件主体，忽略单独扫到的 .lrc 文件
			if utils.MatchAudio(item.Name()) {
				fPath := filepath.Join(fullDirPath, item.Name())
				// 检查同名歌词文件是否存在
				extLrc := filepath.Ext(item.Name())
				// 替换后缀： xxx.mp3 -> xxx.lrc
				lrcName := strings.TrimSuffix(item.Name(), extLrc) + ".lrc"
				lrcPath := filepath.Join(fullDirPath, lrcName)
				hasLocalLrc := false // 标记本地是否已有歌词

				if _, err := os.Stat(lrcPath); err == nil {
					hasLocalLrc = true
				}

				songInfo, err := utils.ParseAudioWithFFprobe(fPath, !hasLocalLrc)
				if err != nil {
					log.Printf("解析音频文件失败 %s: %v\n", fPath, err)
					continue
				}
				if hasLocalLrc {
					songInfo.LrcPath = lrcPath
				}

				result <- *songInfo
			}
		}
	}
}

//  music/周杰伦-魔羯座/xxx.mp3
//  music/周杰伦-魔羯座/xxx.lrc

// 1.根据文件信息解析出对应的源数据信息
// 2.如果对应的文件有歌词文件就不生成歌词文件，而是获取歌词文件的路径
// 3.如果对应的文件没有歌词文件就在源数据里面获取，如果元数据里面有歌词信息，并生成歌词文件，并存入到文件里面，同时获取路径，如果源数据没有歌词，就不做任何处理

// 扫描状态
func (s *ScanSerImpl) GetScanStatus() bool {
	s.mu.RLock()

	defer s.mu.RUnlock()
	return s.IsScanning
}

// 扫描操作
func (s *ScanSerImpl) Scan(targetDirs []string) {
	s.mu.Lock()

	if s.IsScanning {
		s.mu.Unlock()
		return
	}
	s.IsScanning = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		s.IsScanning = false
		s.mu.Unlock()
	}()

	var wg sync.WaitGroup

	musicFileChan := make(chan model.SongDB, 100)

	rootDir, err := os.Getwd()
	if err != nil {
		log.Printf("读取根目录失败 %v\n", err)
		return
	}
	for _, dir := range targetDirs {
		wg.Add(1)
		// 判断是否是绝对路径
		if filepath.IsAbs(dir) {
			rootDir = ""
		}
		go ScanDir(rootDir, dir, musicFileChan, &wg)
	}

	go func() {
		wg.Wait()
		close(musicFileChan)
	}()

	var total int

	var fileArr []model.SongDB

	// 主协程：只负责从通道接收数据，不做耗时的数据库 IO
	for fileInfo := range musicFileChan {

		fileArr = append(fileArr, fileInfo)

		total++

	}
	if len(fileArr) > 0 {

		err := s.ScanRepo.ScanDataBatchSave(fileArr)

		if err != nil {
			log.Printf("批量入库失败: %v\n", err)
		} else {
			log.Println("批量入库成功！")
		}

	} else {
		log.Println("未扫描到任何音乐文件")
	}

	log.Printf("总共%d文件", total)
}
