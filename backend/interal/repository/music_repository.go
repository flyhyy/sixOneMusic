package repository

import (
	"errors"

	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type MusicRepoImpl struct {
	DB *gorm.DB
}

type MusicRepoFace interface {
	MusicQueryAll(page int, pageSize int, userId uint) ([]model.MusicBase, error)
	MusicCollect(songId uint, userId uint, is_collect bool) error
	GetMusicCollect(userId uint) ([]model.MusicBase, error)
}

func NewMusicRepo(db *gorm.DB) MusicRepoFace {
	return &MusicRepoImpl{DB: db}
}

func (m *MusicRepoImpl) MusicQueryAll(page int, pageSize int, userId uint) ([]model.MusicBase, error) {
	var data []model.MusicBase

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	sqlStr := `
		SELECT    song.*, sc.id IS NOT NULL
		 AS
		  is_collect FROM song_dbs AS song 
		  LEFT JOIN song_collects_dbs AS sc ON song.id = sc.song_id AND sc.user_id = ?
		  LIMIT ? OFFSET ?
	`

	err := m.DB.Raw(sqlStr, userId, pageSize, offset).Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *MusicRepoImpl) MusicCollect(songId uint, userId uint, is_collect bool) error {
	if is_collect {
		var count int64
		err := m.DB.Model(&model.SongCollectsDB{}).Where("song_id = ? AND user_id = ?", songId, userId).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("该歌曲已收藏，请勿重复添加")
		}
		data := model.SongCollectsDB{
			SongID: songId,
			UserID: userId,
		}
		err = m.DB.Create(&data).Error
		if err != nil {
			return err
		}

	} else {
		res := m.DB.Where("user_id = ? AND song_id = ?", userId, songId).Unscoped().Delete(&model.SongCollectsDB{})

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected > 0 {
			return nil
		}

	}
	return nil
}

func (m *MusicRepoImpl) GetMusicCollect(userId uint) ([]model.MusicBase, error) {
	var data []model.MusicBase

	sqlStr := `
		SELECT song.*, 1 AS is_collect  FROM song_dbs  AS song  INNER JOIN song_collects_dbs AS sc ON sc.song_id = song.id AND sc.user_id = ?
	`

	err := m.DB.Raw(sqlStr, userId).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
