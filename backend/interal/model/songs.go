package model

import "gorm.io/gorm"

type SongDB struct {
	gorm.Model
	Title     string    `gorm:"type:varchar(255);not null;index"` // 歌曲名  唯一索引，不能为空
	Singer    string    `gorm:"type:varchar(255);not null;index"` // 歌手
	StyleName string    `gorm:"type:varchar(50);index"`           // 曲风
	AlbumID   uint      // 物理外键，存在数据库里的纯数字
	AlbumInfo AlbumDB   `gorm:"foreignKey:AlbumID"` // 关联对象，用于 Preload 预加载查询
	AlbumName string    `gorm:"type:varchar(100)"`
	FilePath  string    `gorm:"type:varchar(500);not null;uniqueIndex"` // 在nas中的路径
	LrcPath   string    `gorm:"type:varchar(500);"`                     // 歌词路径
	CoverPath string    `gorm:"type:varchar(500);"`
	Duration  int       `gorm:"comment:歌曲时长(秒)"` // 时长
	Size      int64     `gorm:"comment:文件大小"`    // 文件大小
	Styles    []StyleDB `gorm:"many2many:song_styles"`
}

type SongCollectsDB struct {
	gorm.Model
	SongID   uint   `gorm:"not null;uniqueIndex:idx_user_song"`
	UserID   uint   `gorm:"not null;uniqueIndex:idx_user_song"`
	SongInfo SongDB `gorm:"foreignKey:SongID"`
}
