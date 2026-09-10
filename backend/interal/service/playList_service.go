package service

import (
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
)

type PlayListSerImpl struct {
	PlayListRepo repository.PlayListFace
}

type PlaListSerFace interface {
	Add(name string, userId uint) error
	Query() ([]model.PlayListBase, error)
	Del(playListId uint) error
	Update(data model.PlayListBase) error
	SaveOrUpdatePlayListSong(data model.PlayListAddSongItemRequest) error
	GetPlayListSongs(playListId uint) ([]model.SongDB, error)
}

func NewPLayListSer(repo repository.PlayListFace) PlaListSerFace {
	return &PlayListSerImpl{
		PlayListRepo: repo,
	}
}

func (p *PlayListSerImpl) Add(name string, userId uint) error {
	return p.PlayListRepo.PlayListAdd(name, userId)
}

func (p *PlayListSerImpl) Query() ([]model.PlayListBase, error) {
	return p.PlayListRepo.GetPlayList()
}

func (p *PlayListSerImpl) Del(playListId uint) error {
	return p.PlayListRepo.DelPlayList(playListId)
}

func (p *PlayListSerImpl) Update(data model.PlayListBase) error {
	return p.PlayListRepo.UpdatePlayList(data)
}

func (p *PlayListSerImpl) SaveOrUpdatePlayListSong(data model.PlayListAddSongItemRequest) error {
	return p.PlayListRepo.SaveOrUpdateBatch(data)
}

func (p *PlayListSerImpl) GetPlayListSongs(playListId uint) ([]model.SongDB, error) {
	return p.PlayListRepo.GetPlayListSongs(playListId)
}
