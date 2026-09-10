package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitSingerRouter(apiGroup *gin.RouterGroup, ctr *control.SingerCtrImpl) {
	apiGroup.Use(middleware.JWTAuth())
	group := apiGroup.Group("/singer")
	{
		group.GET("/list", ctr.GetSingerListHandler)
		group.GET("/song/:singerName", ctr.GetSingerSongListHandler)
	}
}
