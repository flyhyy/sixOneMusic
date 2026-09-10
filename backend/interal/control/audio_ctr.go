package control

import (
	"log"
	"os"
	"strconv"

	"fnMusicServe/interal/config"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type AudioCtrImpl struct {
	PlaySer service.AudioSerFace
}

type AudioCtrFace interface {
	AudioPlayHandler(*gin.Context)
}

func NewAudioCtr(ser service.AudioSerFace) *AudioCtrImpl {
	return &AudioCtrImpl{PlaySer: ser}
}

func (m *AudioCtrImpl) AudioPlayHandler(g_ctx *gin.Context) {
	songIdUint64, err := strconv.ParseUint(g_ctx.Param("id"), 10, 32)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	songId := uint(songIdUint64)

	audioPath, err := m.PlaySer.GetAudioPath(songId)
	if err != nil {
		// utils.Error(g_ctx, config.CodeError, "无效的歌曲ID")
		g_ctx.String(400, err.Error())
		return
	}

	// 检查文件是否存在
	_, err = os.Stat(audioPath)

	if os.IsNotExist(err) {

		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}

	g_ctx.File(audioPath)
}

func (m *AudioCtrImpl) AudioLrcHandler(g_ctx *gin.Context) {
	songIdUint64, err := strconv.ParseUint(g_ctx.Param("id"), 10, 32)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	songId := uint(songIdUint64)

	lrcPath, err := m.PlaySer.GetAudioLrc(songId)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, "无效的歌曲ID")
		return
	}

	// C:\Users\47211\Desktop\OpenCode\fnMusicServe\music\周杰伦-七里香\01 - 我的地盤.flac
	// C:\Users\47211\Desktop\OpenCode\fnMusicServe\music\周杰伦-七里香\01 - 我的地盤.lrc
	// 检查文件是否存在
	_, err = os.Stat(lrcPath)

	if os.IsNotExist(err) {
		log.Printf("======= err.Error()====%v", err.Error())
		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}
	g_ctx.File(lrcPath)
}
