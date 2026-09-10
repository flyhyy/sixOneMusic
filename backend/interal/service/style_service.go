package service

import (
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
)

type StyleSerImpl struct {
	StyleRepo repository.StyleRepoFace
}

type StyleSerFace interface {
	StyleList() ([]model.StyleListResoponse, error)
	StyleQuerySong(userId uint, styleId uint) ([]model.MusicBase, error)
}

func NewStyleSer(repo repository.StyleRepoFace) StyleSerFace {
	return &StyleSerImpl{StyleRepo: repo}
}

func (s *StyleSerImpl) StyleList() ([]model.StyleListResoponse, error) {
	return s.StyleRepo.StyleList()
}

func (s *StyleSerImpl) StyleQuerySong(userId uint, styleId uint) ([]model.MusicBase, error) {
	return s.StyleRepo.StyleQuerySong(userId, styleId)
}
