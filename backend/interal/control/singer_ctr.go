package control

import (
	"fnMusicServe/interal/config"
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type SingerCtrImpl struct {
	SingerSer service.SingerSerFace
}

func NewSingerCtr(ser service.SingerSerFace) *SingerCtrImpl {
	return &SingerCtrImpl{SingerSer: ser}
}

func (s *SingerCtrImpl) GetSingerListHandler(g_ctx *gin.Context) {
	var list []model.SingerListRespose

	arr, err := s.SingerSer.SingerList()
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	for _, item := range arr {
		list = append(list, model.SingerListRespose{
			Name:          item.Name,
			TotalDuration: item.TotalDuration,
			TotalNum:      item.TotalNum,
		})
	}

	utils.Success(g_ctx, list)
}

func (s *SingerCtrImpl) GetSingerSongListHandler(g_ctx *gin.Context) {
	name := g_ctx.Param("singerName")

	if name == "" {
		utils.Error(g_ctx, config.CodeError, "歌手名称为空")
		return
	}
	userId, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, "用户信息获取失败")

		return

	}

	list, err := s.SingerSer.QuerySingerSongList(name, userId)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, list)
}
