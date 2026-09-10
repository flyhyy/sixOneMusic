package model

import "gorm.io/gorm"

// 专辑结构体
type AlbumDB struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);uniqueIndex;NOT NULL;"`

	Singer         string `gorm:"type:varchar(50)"`
	TotoalDuration int
	TotalNum       int64
	ReleaseDate    string
	CoverPath      string   `gorm:"type:varchar(255)"`
	Songs          []SongDB `gorm:"foreignKey:AlbumID"`
}

type AlbumListRespose struct {
	Name           string `json:"name"`
	Singer         string `json:"singer"`
	TotoalDuration int    `json:"totoal_duration"`
	TotalNum       int64  `json:"total_num"`
	CoverPath      string `json:"cover_path"`
}
