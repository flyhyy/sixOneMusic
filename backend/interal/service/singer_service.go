package service

import (
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
)

type SingerSerImpl struct {
	SingerRepo repository.SingerRepoFace
}

type SingerSerFace interface {
	SingerList() ([]model.SingerDB, error)
	QuerySingerSongList(singerName string, userId uint) ([]model.MusicBase, error)
}

func NewSingerSer(repo repository.SingerRepoFace) SingerSerFace {
	return &SingerSerImpl{SingerRepo: repo}
}

func (s *SingerSerImpl) SingerList() ([]model.SingerDB, error) {
	return s.SingerRepo.SingerList()
}

func (s *SingerSerImpl) QuerySingerSongList(singerName string, userId uint) ([]model.MusicBase, error) {
	return s.SingerRepo.QuerySingerSongList(singerName, userId)
}
