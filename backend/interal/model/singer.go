package model

import "gorm.io/gorm"

type SingerDB struct {
	gorm.Model
	Name          string `gorm:"type:varchar(50);NOT NULL; uniqueIndex"`
	TotalDuration int    `gorm:"NOT NULL;"`
	TotalNum      int64  `gorm:"NOT NULL;"`
}

type SingerListRespose struct {
	Name          string `json:"name"`
	TotalDuration int    `json:"total_duration"`
	TotalNum      int64  `json:"total_num"`
}

type QuerySingerSongListRequset struct {
	Name string `json:"name"`
}
