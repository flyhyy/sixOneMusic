package repository

import (
	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type StyleRepoImpl struct {
	DB *gorm.DB
}
type StyleRepoFace interface {
	StyleList() ([]model.StyleListResoponse, error)
	StyleQuerySong(userId uint, styleId uint) ([]model.MusicBase, error)
}

func NewStyleRepo(db *gorm.DB) StyleRepoFace {
	return &StyleRepoImpl{DB: db}
}

func (s *StyleRepoImpl) StyleDataBatchSave(songs []model.SongDB) {
}

func (s *StyleRepoImpl) StyleList() ([]model.StyleListResoponse, error) {
	var list []model.StyleListResoponse
	res := s.DB.Find(&model.StyleDB{}).Scan(&list)

	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

func (s *StyleRepoImpl) StyleQuerySong(userId uint, styleId uint) ([]model.MusicBase, error) {
	var list []model.MusicBase

	styleDB := model.StyleDB{}
	styleDB.ID = styleId
	var songDB []model.SongDB

	err := s.DB.Model(&styleDB).Association("Song").Find(&songDB)
	if err != nil {
		return nil, err
	}
	var collectSongsIds []uint
	// 2. 查出当前用户收藏的所有 歌曲ID（Pluck 专门用来只查某一列）
	s.DB.Model(&model.SongCollectsDB{}).Where("user_id = ?", userId).Pluck("song_id", &collectSongsIds)

	collectMap := make(map[uint]bool)

	for _, id := range collectSongsIds {
		collectMap[id] = true
	}

	for _, item := range songDB {
		list = append(list, model.MusicBase{
			Title:     item.Title,
			Singer:    item.Singer,
			Duration:  item.Duration,
			Size:      item.Size,
			ID:        item.ID,
			CoverPath: item.CoverPath,
			IsCollect: collectMap[item.ID],
			AlbumName: item.AlbumName,
		})
	}
	if list == nil {
		list = make([]model.MusicBase, 0)
	}
	return list, nil
}
