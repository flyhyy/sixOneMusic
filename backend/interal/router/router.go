package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/repository"
	"fnMusicServe/interal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// ================= 1. 初始化所有的 Repository 层 =================
	userRepo := repository.NewUserRepo(db)
	folderRepo := repository.NewFolderRepo(db)
	scanRepo := repository.NewScanRepo(db)
	musicRepo := repository.NewMusicRepo(db)
	audioRepo := repository.NewAudioRepo(db)
	playListRepo := repository.NewPLayListRepo(db)
	styleRepo := repository.NewStyleRepo(db)
	singerRepo := repository.NewSingerRepo(db)
	albumRepo := repository.NewAlbumRepo(db)
	// ================= 2. 初始化所有的 Service 层 =================
	authSer := service.NewAuthSer(userRepo)
	folderSer := service.NewFolderSer(folderRepo)
	scanSer := service.NewScanSer(scanRepo)
	musicSer := service.NewMusicSer(musicRepo)
	audioSer := service.NewAudioSer(audioRepo)
	playListSer := service.NewPLayListSer(playListRepo)
	styleSer := service.NewStyleSer(styleRepo)
	singerSer := service.NewSingerSer(singerRepo)
	albumSer := service.NewAlbumSer(albumRepo)
	// ================= 3. 初始化所有的 Control 层 =================
	authCtr := control.NewAuthControl(authSer)
	folderCtr := control.NewFolderCtr(folderSer)
	scanCtr := control.NewScanCtr(scanSer, folderSer)
	musicCtr := control.NewMusicCtr(musicSer)
	audioCtr := control.NewAudioCtr(audioSer)
	playListCtr := control.NewPLayListCtr(playListSer)
	styleCtr := control.NewStyleCtr(styleSer)
	singerCtr := control.NewSingerCtr(singerSer)
	albumCtr := control.NewAlbumCtr(albumSer)

	AuthRouter(api, authCtr)
	InitFolderRouter(api, folderCtr)
	InitScanRouter(api, scanCtr)
	InitMusicRouter(api, musicCtr)
	InitAudioRouter(api, audioCtr)
	InitPlayListRouter(api, playListCtr)
	InitStyleRouter(api, styleCtr)
	InitSingerRouter(api, singerCtr)
	InitAlbumRouter(api, albumCtr)
	return r
}
