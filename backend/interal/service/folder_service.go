package service

import (
	"errors"
	"os"

	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
)

type FolderSceImpl struct {
	FolderPathRepo repository.FolderPathRepoFace
}

type FolderSceFace interface {
	Query() ([]model.FolderResponse, error)
	Write([]model.FolderWriteItemRequest) error
	Delete(item model.FolderPathDeleteItemRequest) error
}

func NewFolderSer(repo repository.FolderPathRepoFace) FolderSceFace {
	return &FolderSceImpl{FolderPathRepo: repo}
}

func (f *FolderSceImpl) Query() ([]model.FolderResponse, error) {
	arr, err := f.FolderPathRepo.FolderQuery()

	var data []model.FolderResponse

	for _, item := range arr {
		data = append(data, model.FolderResponse{ID: item.ID, Path: item.Path})
	}

	if err != nil {
		return []model.FolderResponse{}, err
	}
	return data, nil
}

func (f *FolderSceImpl) Write(paths []model.FolderWriteItemRequest) error {
	// 增加物理路径校验
	for _, item := range paths {

		info, err := os.Stat(item.Path)
		if err != nil {
			return errors.New("路径不存在" + err.Error())
		}
		if !info.IsDir() {
			return errors.New("提供的路径不是一个文件夹:" + item.Path)
		}

	}

	err := f.FolderPathRepo.FolderBatchSave(paths)
	if err != nil {
		return err
	}

	return nil
}

func (f *FolderSceImpl) Delete(item model.FolderPathDeleteItemRequest) error {
	err := f.FolderPathRepo.FolderDel(item.ID)
	return err
}
