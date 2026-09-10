package service

import (
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
)

type AlbumSerImpl struct {
	AblumRepo repository.AlbumRepoFace
}

type AlbumSerFace interface {
	AlbumList() ([]model.AlbumDB, error)
	AlbumSongList(alnumName string, userId uint) ([]model.MusicBase, error)
}

func NewAlbumSer(repo repository.AlbumRepoFace) AlbumSerFace {
	return &AlbumSerImpl{AblumRepo: repo}
}

func (a *AlbumSerImpl) AlbumList() ([]model.AlbumDB, error) {
	return a.AblumRepo.AlbumList()
}

func (a *AlbumSerImpl) AlbumSongList(alnumName string, userId uint) ([]model.MusicBase, error) {
	return a.AblumRepo.AlbumSongList(alnumName, userId)
}
