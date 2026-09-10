package repository

import (
	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type AudioRepoImpl struct {
	DB *gorm.DB
}

type AudioRepoFace interface {
	AudioItemPath(id uint) (string, error)
	AudioLrc(id uint) (string, error)
}

func NewAudioRepo(db *gorm.DB) AudioRepoFace {
	return &AudioRepoImpl{DB: db}
}

func (m *AudioRepoImpl) AudioItemPath(id uint) (string, error) {
	var song model.SongDB
	err := m.DB.Select("file_path").First(&song, id).Error
	if err != nil {
		return "", err
	}
	return song.FilePath, nil
}

func (m *AudioRepoImpl) AudioLrc(id uint) (string, error) {
	var song model.SongDB
	err := m.DB.Select("lrc_path").First(&song, id).Error
	if err != nil {
		return "", err
	}
	return song.LrcPath, nil
}
