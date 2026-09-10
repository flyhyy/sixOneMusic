package model

import "gorm.io/gorm"

// 请求参数
type RequestLogin struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 返回参数

type ResponseLogin struct {
	Token string `json:"token"`
}

// 定义结构体
type UserDB struct {
	gorm.Model
	UserName  string       `gorm:"uniqueIndex;not null;type:varchar(50)"`
	Password  string       `gorm:"not null"`
	Playlists []PlayListDB `gorm:"foreignKey:UserID"`
}
