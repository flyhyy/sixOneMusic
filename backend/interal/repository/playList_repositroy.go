package repository

import (
	"errors"

	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type PlayListRepoImpl struct {
	DB *gorm.DB
}

type PlayListFace interface {
	PlayListAdd(name string, userId uint) error
	GetPlayList() ([]model.PlayListBase, error)
	DelPlayList(playId uint) error
	UpdatePlayList(data model.PlayListBase) error
	SaveOrUpdateBatch(data model.PlayListAddSongItemRequest) error
	GetPlayListSongs(playListId uint) ([]model.SongDB, error)
}

func NewPLayListRepo(db *gorm.DB) PlayListFace {
	return &PlayListRepoImpl{DB: db}
}

func (p *PlayListRepoImpl) PlayListAdd(name string, userId uint) error {
	data := model.PlayListDB{
		Name:   name,
		UserID: userId,
	}

	err := p.DB.Model(&model.PlayListDB{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PlayListRepoImpl) GetPlayList() ([]model.PlayListBase, error) {
	// var DBlist []model.PlayListDB
	var list []model.PlayListBase

	// err := p.DB.Find(&DBlist).Error
	// if err != nil {
	// 	return list, err
	// }

	// for _, item := range DBlist {
	// 	list = append(list, model.PlayListBase{
	// 		Name: item.Name,
	// 		ID:   item.ID,
	// 	})
	// }
	// return list, nil

	err := p.DB.Model(&model.PlayListDB{}).Select("id", "name").Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (p *PlayListRepoImpl) DelPlayList(playListId uint) error {
	if playListId == 0 {
		return errors.New("无效的删除ID")
	}

	playList := model.PlayListDB{}
	playList.ID = playListId

	result := p.DB.Select("Songs").Unscoped().Delete(&playList)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("记录不存在或者已经被删除")
	}
	return nil
}

func (p *PlayListRepoImpl) UpdatePlayList(data model.PlayListBase) error {
	if data.ID == 0 {
		return errors.New("无效的ID")
	}
	if data.Name == "" {
		return errors.New("名称为空")
	}

	res := p.DB.Model(&model.PlayListDB{}).Where("id = ? AND user_id = ?", data.ID, data.UserID).Update("name", data.Name)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("歌单不存在或您无权修改该歌单")
	}

	return nil
}

func (p *PlayListRepoImpl) SaveOrUpdateBatch(data model.PlayListAddSongItemRequest) error {
	if data.ID == 0 {
		return errors.New("歌单ID为空")
	}
	if data.SongID == 0 {
		return errors.New("歌曲ID为空")
	}

	var playList model.PlayListDB

	err := p.DB.First(&playList, data.ID).Error
	if err != nil {
		return errors.New("未查询到对应歌单")
	}

	var existSongs []model.SongDB

	err = p.DB.Model(&playList).Association("Songs").Find(&existSongs, data.SongID)
	if err != nil {
		return err
	}
	if len(existSongs) > 0 {
		return errors.New("该歌曲已存在歌单中,请勿重复添加")
	}

	song := model.SongDB{}
	song.ID = data.SongID

	err = p.DB.Model(&playList).Association("Songs").Append(&song)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlayListRepoImpl) GetPlayListSongs(playListId uint) ([]model.SongDB, error) {
	if playListId == 0 {
		return nil, errors.New("歌单ID无效")
	}

	playList := model.PlayListDB{}
	playList.ID = playListId

	var songs []model.SongDB

	err := p.DB.Model(&playList).Association("Songs").Find(&songs)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
