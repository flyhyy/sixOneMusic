package repository

import (
	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type SingerRepoImpl struct {
	DB *gorm.DB
}

type SingerRepoFace interface {
	SingerList() ([]model.SingerDB, error)
	QuerySingerSongList(singerName string, userId uint) ([]model.MusicBase, error)
}

func NewSingerRepo(db *gorm.DB) SingerRepoFace {
	return &SingerRepoImpl{DB: db}
}

func (s *SingerRepoImpl) SingerList() ([]model.SingerDB, error) {
	var list []model.SingerDB

	res := s.DB.Find(&list)
	if res.Error != nil {
		return make([]model.SingerDB, 0), res.Error
	}
	return list, nil
}

// 查询歌手对应的歌曲
func (s *SingerRepoImpl) QuerySingerSongList(singerName string, userId uint) ([]model.MusicBase, error) {
	var list []model.MusicBase

	var songList []model.SongDB

	var collectSongsIds []uint
	s.DB.Model(&model.SongCollectsDB{}).Where("user_id = ?", userId).Pluck("song_id", &collectSongsIds)

	collectMap := make(map[uint]bool)

	for _, id := range collectSongsIds {
		collectMap[id] = true
	}

	s.DB.Where("singer = ?", singerName).Find(&songList)

	for _, item := range songList {
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
