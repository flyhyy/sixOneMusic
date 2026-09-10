package control

import (
	"strconv"

	"fnMusicServe/interal/config"
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type PlayListCtrImpl struct {
	PlayListSer service.PlaListSerFace
}

type PlayListCtrFace interface {
	Add(g_ctx *gin.Context)
}

func NewPLayListCtr(ser service.PlaListSerFace) *PlayListCtrImpl {
	return &PlayListCtrImpl{
		PlayListSer: ser,
	}
}

func (p *PlayListCtrImpl) Add(g_ctx *gin.Context) {
	var req model.PlayListAddRequest

	err := g_ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	userId, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	err = p.PlayListSer.Add(req.Name, userId)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}
	utils.Success(g_ctx, nil)
}

func (p *PlayListCtrImpl) Query(g_ctx *gin.Context) {
	arr, err := p.PlayListSer.Query()
	if err != nil {

		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, arr)
}

func (p *PlayListCtrImpl) Del(g_ctx *gin.Context) {
	id := g_ctx.Param("id")

	if id == "" {
		utils.Error(g_ctx, config.CodeError, "歌单ID为空")
		return
	}
	playListId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	err = p.PlayListSer.Del(uint(playListId))
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}
	utils.Success(g_ctx, nil)
}

func (p *PlayListCtrImpl) Update(g_ctx *gin.Context) {
	var req model.PlayListUpdateRequest

	err := g_ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	uid, err := utils.GetUserID(g_ctx)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	data := model.PlayListBase{
		ID:     req.ID,
		Name:   req.Name,
		UserID: uid,
	}

	err = p.PlayListSer.Update(data)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, nil)
}

func (p *PlayListCtrImpl) SaveOrUpdatePlayListSongHandler(g_ctx *gin.Context) {
	var req model.PlayListAddSongItemRequest

	err := g_ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	err = p.PlayListSer.SaveOrUpdatePlayListSong(req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	utils.Success(g_ctx, nil)
}

func (p *PlayListCtrImpl) GetPlayListSongsHandler(g_ctx *gin.Context) {
	id := g_ctx.Param("id")

	if id == "" {
		utils.Error(g_ctx, config.CodeError, "歌单ID为空")
		return
	}
	playListId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	songs, err := p.PlayListSer.GetPlayListSongs(uint(playListId))
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	var responseSongs []model.MusicBase

	for _, item := range songs {
		responseSongs = append(responseSongs, model.MusicBase{
			ID:        item.ID,
			Title:     item.Title,
			Singer:    item.Singer,
			Duration:  item.Duration,
			Size:      item.Size,
			CoverPath: item.CoverPath,
			// IsCollect: *item.IsCollect,
			AlbumName: item.AlbumName,
		})
	}
	utils.Success(g_ctx, responseSongs)
}
