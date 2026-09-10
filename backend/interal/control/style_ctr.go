package control

import (
	"fnMusicServe/interal/config"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type StyleCtrImpl struct {
	StyleSer service.StyleSerFace
}

func NewStyleCtr(ser service.StyleSerFace) *StyleCtrImpl {
	return &StyleCtrImpl{StyleSer: ser}
}

func (s *StyleCtrImpl) StyleQueryHandle(g_ctx *gin.Context) {
	data, err := s.StyleSer.StyleList()
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	utils.Success(g_ctx, data)
}

func (s *StyleCtrImpl) StyleQuerySongHandle(g_ctx *gin.Context) {
	styleId, err := utils.StrToUnit(g_ctx.Param("styleId"))
	if err != nil {

		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	userId, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	list, err := s.StyleSer.StyleQuerySong(userId, styleId)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	utils.Success(g_ctx, list)
}
