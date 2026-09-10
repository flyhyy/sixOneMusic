package repository

import (
	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type AlbumRepoImpl struct {
	DB *gorm.DB
}

type AlbumRepoFace interface {
	AlbumList() ([]model.AlbumDB, error)
	AlbumSongList(alnumName string, userId uint) ([]model.MusicBase, error)
}

func NewAlbumRepo(db *gorm.DB) AlbumRepoFace {
	return &AlbumRepoImpl{DB: db}
}

func (a *AlbumRepoImpl) AlbumList() ([]model.AlbumDB, error) {
	var list []model.AlbumDB

	res := a.DB.Find(&list)

	if res.Error != nil {
		return make([]model.AlbumDB, 0), res.Error
	}
	return list, nil
}

func (a *AlbumRepoImpl) AlbumSongList(alnumName string, userId uint) ([]model.MusicBase, error) {
	var list []model.MusicBase
	var songList []model.SongDB

	var collectSongsIds []uint
	collectMap := make(map[uint]bool)

	a.DB.Model(&model.SongCollectsDB{}).Where("user_id = ?", userId).Pluck("song_id", &collectSongsIds)

	if len(collectSongsIds) > 0 {
		for _, id := range collectSongsIds {
			collectMap[id] = true
		}
	}

	a.DB.Where("album_name = ?", alnumName).Find(&songList)

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
