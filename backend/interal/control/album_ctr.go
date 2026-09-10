package control

import (
	"fnMusicServe/interal/config"
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type AlbumCtrImpl struct {
	AlbumSer service.AlbumSerFace
}

func NewAlbumCtr(ser service.AlbumSerFace) *AlbumCtrImpl {
	return &AlbumCtrImpl{AlbumSer: ser}
}

func (a *AlbumCtrImpl) GetAlbumListHandle(g_ctx *gin.Context) {
	var data []model.AlbumListRespose

	list, err := a.AlbumSer.AlbumList()
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	for _, item := range list {
		data = append(data, model.AlbumListRespose{
			Name:           item.Title,
			Singer:         item.Singer,
			TotoalDuration: item.TotoalDuration,
			TotalNum:       item.TotalNum,
			CoverPath:      item.CoverPath,
		})
	}

	utils.Success(g_ctx, data)
}

func (a *AlbumCtrImpl) GetAlbumSongListHandle(g_ctx *gin.Context) {
	albumName := g_ctx.Param("albumName")

	if albumName == "" {
		utils.Error(g_ctx, config.CodeError, "专辑名称为空")
		return
	}

	userId, err := utils.GetUserID(g_ctx)
	if err != nil {

		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}

	list, err := a.AlbumSer.AlbumSongList(albumName, userId)
	if err != nil {

		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, list)
}
