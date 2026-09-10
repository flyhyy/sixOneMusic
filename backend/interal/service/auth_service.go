package service

import (
	"errors"
	"log"

	"fnMusicServe/interal/model"
	"fnMusicServe/interal/repository"
	"fnMusicServe/interal/utils"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepoFace
}

type UserServiceFace interface {
	Register(usename string, password string) error
	Login(username string, password string) (*model.ResponseLogin, error)
}

func NewAuthSer(repo repository.UserRepoFace) UserServiceFace {
	return &UserServiceImpl{UserRepo: repo}
}

// 注册服务
func (u *UserServiceImpl) Register(username string, password string) error {
	isExist, err := u.UserRepo.CheckUserExist(username)
	if err != nil {
		return err
	}

	if isExist {
		log.Printf("注册用户存在:%v", err)
		return errors.New("注册用户存在")
	}
	password, err = utils.HashPassword(password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	user := &model.UserDB{
		UserName: username,
		Password: password,
	}
	return u.UserRepo.CreateUser(user)
}

// 登录服务
func (u *UserServiceImpl) Login(username string, password string) (*model.ResponseLogin, error) {
	user, err := u.UserRepo.SelectUserInfo(username)
	log.Println(err)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	err = utils.PasswordDiff(user.Password, password)
	if err != nil {
		return nil, errors.New("密码错误")
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return nil, err
	}

	res := &model.ResponseLogin{
		Token: token,
	}

	return res, nil
}
