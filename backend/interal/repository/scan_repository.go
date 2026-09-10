package repository

import (
	"errors"
	"log"

	"fnMusicServe/interal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScanImpl struct {
	DB *gorm.DB
}

type ScanRepoFace interface {
	SacningData() ([]model.FolderScanList, error)
	ScanDataBatchSave(data []model.SongDB) error
}

func NewScanRepo(db *gorm.DB) ScanRepoFace {
	return &ScanImpl{DB: db}
}

// 扫描数据
func (s *ScanImpl) SacningData() ([]model.FolderScanList, error) {
	var folders []model.FolderScanList

	err := s.DB.Model(&model.FolderDB{}).Select("path").Find(&folders).Error
	if err != nil {
		return nil, err
	}
	return folders, nil
}

// 存入扫描数据

type SingerStats struct {
	TotalDuration int
	TotalNum      int64
}

type AlbumStats struct {
	Singer         string
	TotoalDuration int
	TotalNum       int64
	CoverPath      string
}

func (s *ScanImpl) ScanDataBatchSave(songs []model.SongDB) error {
	if len(songs) == 0 {
		return errors.New("歌曲列表为空")
	}

	almIdMap := make(map[string]uint)
	var styleRecords []model.StyleDB

	almMap := make(map[string]*AlbumStats)

	singerMap := make(map[string]*SingerStats)

	for _, item := range songs {

		var style model.StyleDB
		s.DB.Where(model.StyleDB{Name: item.StyleName}).FirstOrCreate(&style)
		styleRecords = append(styleRecords, style)

		if item.AlbumName != "" {
			if almMap[item.AlbumName] == nil {
				almMap[item.AlbumName] = &AlbumStats{}
			}
			almMap[item.AlbumName].Singer = item.Singer
			almMap[item.AlbumName].TotalNum++
			almMap[item.AlbumName].TotoalDuration += item.Duration
			almMap[item.AlbumName].CoverPath = item.CoverPath

		}

		if item.Singer != "" {
			if singerMap[item.Singer] == nil {
				singerMap[item.Singer] = &SingerStats{}
			}
			singerMap[item.Singer].TotalDuration += item.Duration
			singerMap[item.Singer].TotalNum++
		}

	}

	if len(almMap) > 0 {

		var albumBatch []model.AlbumDB
		var albumNames []string

		for key, item := range almMap {
			albumNames = append(albumNames, key)
			albumBatch = append(albumBatch, model.AlbumDB{
				Title:          key,
				Singer:         item.Singer,
				TotoalDuration: item.TotoalDuration,
				TotalNum:       item.TotalNum,
			})
		}

		err := s.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "Title"}},
			UpdateAll: true,
		}).CreateInBatches(&albumBatch, 100).Error
		if err != nil {
			log.Printf("专辑批量入库/更新失败: %v", err)
		}

		var saveAlumn []model.AlbumDB
		s.DB.Where("title IN ?", albumNames).Find(&saveAlumn)

		for _, item := range saveAlumn {
			almIdMap[item.Title] = item.ID
		}

	}

	if len(singerMap) > 0 {

		var singerBatch []model.SingerDB

		for k, item := range singerMap {
			singerBatch = append(singerBatch, model.SingerDB{
				Name:          k,
				TotalDuration: item.TotalDuration,
				TotalNum:      item.TotalNum,
			})
		}

		err := s.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "Name"}},
			UpdateAll: true,
		}).CreateInBatches(&singerBatch, 100).Error
		if err != nil {
			log.Printf("歌手批量入库/更新失败: %v", err)
		}

	}

	for i := range songs {
		if songs[i].AlbumName != "" {
			songs[i].AlbumID = almIdMap[songs[i].AlbumName]
			songs[i].Styles = styleRecords
		}
	}
	log.Printf("========songs===========%+v", songs)

	return s.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "file_path"}},
		UpdateAll: true,
	}).CreateInBatches(&songs, 100).Error
}
