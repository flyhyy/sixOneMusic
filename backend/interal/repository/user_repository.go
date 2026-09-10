package repository

import (
	"fnMusicServe/interal/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

type UserRepoFace interface {
	CheckUserExist(username string) (bool, error)
	CreateUser(user *model.UserDB) error
	SelectUserInfo(username string) (*model.UserDB, error)
}

func NewUserRepo(db *gorm.DB) UserRepoFace {
	return &UserRepositoryImpl{DB: db}
}

// 检查注册用户是否存在
func (u *UserRepositoryImpl) CheckUserExist(username string) (bool, error) {
	var count int64

	err := u.DB.Model(&model.UserDB{}).Where("user_name = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 创建用户
func (u *UserRepositoryImpl) CreateUser(user *model.UserDB) error {
	return u.DB.Create(user).Error
}

// 查询用户信息
func (u *UserRepositoryImpl) SelectUserInfo(username string) (*model.UserDB, error) {
	var user model.UserDB

	result := u.DB.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
