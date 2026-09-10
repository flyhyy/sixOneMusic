package control

import (
	"strconv"

	"fnMusicServe/interal/config"
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type MusicCtrImpl struct {
	MusicSer service.MusicSerFace
}

type MusicCrtFace interface {
	QueryAllHandle(g_ctx *gin.Context)
}

func NewMusicCtr(ser service.MusicSerFace) *MusicCtrImpl {
	return &MusicCtrImpl{MusicSer: ser}
}

func (m *MusicCtrImpl) QueryAllHandle(g_ctx *gin.Context) {
	page, err := strconv.Atoi(g_ctx.Query("page"))
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return

	}

	pageSize, err := strconv.Atoi(g_ctx.Query("pageSize"))
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	userId, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())

		return

	}
	data, err := m.MusicSer.GetMusicListAll(page, pageSize, userId)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	resp := model.MusicAllResponse{
		Data:     data,
		PageSize: pageSize,
		Page:     page,
	}
	utils.Success(g_ctx, resp)
}

// 收藏
func (m *MusicCtrImpl) CollectHandler(g_ctx *gin.Context) {
	var info model.MusicCollectRequest

	err := g_ctx.ShouldBindJSON(&info)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	userId, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())

		return

	}

	err = m.MusicSer.SetCollect(info.SongID, userId, info.IsCollect)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}
	utils.Success(g_ctx, info.SongID)
}

// 获取用户收藏
func (m *MusicCtrImpl) GetUserSongsCollect(g_ctx *gin.Context) {
	userId, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	data, err := m.MusicSer.GetCollectSongs(userId)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return

	}
	utils.Success(g_ctx, data)
}
