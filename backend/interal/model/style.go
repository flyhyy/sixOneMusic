package model

import "gorm.io/gorm"

type StyleDB struct {
	gorm.Model
	Name string   `gorm:"type:varchar(50);uniqueIndex;not null;"`
	Song []SongDB `gorm:"many2many:song_styles;"`
}

type StyleListResoponse struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}
