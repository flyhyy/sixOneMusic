// 歌单
package model

import "gorm.io/gorm"

type PlayListDB struct {
	gorm.Model

	Name   string `gorm:"varchar(50);not null;index;unique"` // 歌单名称
	UserID uint
	Songs  []SongDB `gorm:"many2many:playlist_songs"` // GORM 会自动在数据库建一张叫 `playlist_songs` 的中间表
}

type PlayListBase struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	UserID uint   `json:"-"`
}

type PlayListAddRequest struct {
	Name string `json:"name"`
}

type PlayListUpdateRequest struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

type PlayListDelRequest struct {
	ID uint `json:"id"`
}

type PlayListAddSongItemRequest struct {
	ID     uint `json:"playListId"`
	SongID uint `json:"songId"`
}
