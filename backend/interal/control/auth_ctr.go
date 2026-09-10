package control

import (
	"log"

	"fnMusicServe/interal/config"
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type AuthCtrImpl struct {
	AuthService service.UserServiceFace
}

func NewAuthControl(srv service.UserServiceFace) *AuthCtrImpl {
	return &AuthCtrImpl{AuthService: srv}
}

// 登录
func (ac *AuthCtrImpl) LoginHandler(g_ctx *gin.Context) {
	var req model.RequestLogin
	err := g_ctx.ShouldBindJSON(&req)

	log.Println(req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	tokenInfo, err := ac.AuthService.Login(req.UserName, req.Password)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, tokenInfo)
}

// 注册
func (ac *AuthCtrImpl) RegisterHandler(g_ctx *gin.Context) {
	var req model.RequestLogin

	err := g_ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	err = ac.AuthService.Register(req.UserName, req.Password)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, nil)
}
