package repository

import (
	"errors"
	"fmt"

	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type FolderPathRpoImpl struct {
	DB *gorm.DB
}

type FolderPathRepoFace interface {
	// FolderPathWrite(path string) (bool, error)
	FolderQuery() ([]model.FolderDB, error)
	FolderBatchSave([]model.FolderWriteItemRequest) error
	FolderDel(id uint) error
}

func NewFolderRepo(db *gorm.DB) FolderPathRepoFace {
	return &FolderPathRpoImpl{DB: db}
}

// 路径 批量更新
func (f *FolderPathRpoImpl) FolderBatchSave(list []model.FolderWriteItemRequest) error {
	// 使用 GORM 推荐的闭包事务，自动处理 回滚 和 提交
	return f.DB.Transaction(func(tx *gorm.DB) error {
		var existFolderList []model.FolderDB

		f.DB.Select("path").Find(&existFolderList)

		pathMap := make(map[string]bool)

		for _, item := range existFolderList {
			pathMap[item.Path] = true
		}

		for _, item := range list {
			if item.ID == nil {

				if pathMap[item.Path] {
					return fmt.Errorf("添加失败,路径已经存在:%s", item.Path)
				}

				newFolder := model.FolderDB{
					Path: item.Path,
				}
				if err := tx.Create(&newFolder).Error; err != nil {
					// 返回 err，Transaction 会自动 Rollback
					return err
				}

			} else {
				// / 2. ID 有值，执行更新 (Update)
				// 注意：这里需要解引用 *item.ID 获取实际的值
				err := tx.Model(&model.FolderDB{}).Where("id = ?", *item.ID).Update("path", item.Path).Error
				if err != nil {
					// 返回 err，Transaction 会自动 Rollback
					return err
				}

			}
		}
		// 返回 nil，Transaction 会自动 Commit
		return nil
	})
}

// 路径查询
func (f *FolderPathRpoImpl) FolderQuery() ([]model.FolderDB, error) {
	var folders []model.FolderDB

	err := f.DB.Find(&folders).Error
	if err != nil {
		return folders, err
	}
	return folders, nil
}

// 路径删除
func (f *FolderPathRpoImpl) FolderDel(id uint) error {
	if id == 0 {
		return errors.New("无效的删除ID")
	}

	result := f.DB.Unscoped().Delete(&model.FolderDB{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("记录不存在或已经被删除")
	}

	return nil
}
