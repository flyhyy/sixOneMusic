package service

import (
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
)

type MusicSerImpl struct {
	MusicRepo repository.MusicRepoFace
}

type MusicSerFace interface {
	GetMusicListAll(page int, pageSize int, userId uint) ([]model.MusicBase, error)
	SetCollect(songId uint, userId uint, is_collect bool) error
	GetCollectSongs(userId uint) ([]model.MusicBase, error)
}

func NewMusicSer(repo repository.MusicRepoFace) MusicSerFace {
	return &MusicSerImpl{MusicRepo: repo}
}

func (m *MusicSerImpl) GetMusicListAll(page int, pageSize int, userId uint) ([]model.MusicBase, error) {
	return m.MusicRepo.MusicQueryAll(page, pageSize, userId)
}

func (m *MusicSerImpl) SetCollect(songId uint, userId uint, is_collect bool) error {
	return m.MusicRepo.MusicCollect(songId, userId, is_collect)
}

func (m *MusicSerImpl) GetCollectSongs(userId uint) ([]model.MusicBase, error) {
	return m.MusicRepo.GetMusicCollect(userId)
}
