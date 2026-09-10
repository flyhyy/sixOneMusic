package service

import "fnMusicServe/interal/repository"

type AudioSerImpl struct {
	AudioRepo repository.AudioRepoFace
}

type AudioSerFace interface {
	GetAudioPath(id uint) (string, error)
	GetAudioLrc(id uint) (string, error)
}

func NewAudioSer(repo repository.AudioRepoFace) AudioSerFace {
	return &AudioSerImpl{AudioRepo: repo}
}

func (m *AudioSerImpl) GetAudioPath(id uint) (string, error) {
	return m.AudioRepo.AudioItemPath(id)
}

func (m *AudioSerImpl) GetAudioLrc(id uint) (string, error) {
	return m.AudioRepo.AudioLrc(id)
}
