package model

import "gorm.io/gorm"

type FolderDB struct {
	gorm.Model
	Path string `gorm:"not null;type:varchar(255);uniqueIndex"`
}

// type FolderRequest struct {
// 	Path []string `json:"path"  binding:"required" `
// }

type FolderWriteItemRequest struct {
	ID   *uint  `json:"id"`
	Path string `json:"path" binding:"required"`
}

type FolderResponse struct {
	Path string `json:"path"`
	ID   uint   `json:"id"`
}

// 注意：Query 绑定需要用 `form` 标签，而不是 `json`
type FolderPathDeleteItemRequest struct {
	ID uint `form:"id" binding:"required"`
}

// 列表扫描
type FolderScanList struct {
	Path string
}
