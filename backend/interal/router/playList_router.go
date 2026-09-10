package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitPlayListRouter(apiGroup *gin.RouterGroup, ctr *control.PlayListCtrImpl) {
	group := apiGroup.Group("playList")
	group.Use(middleware.JWTAuth())
	{
		group.POST("/add", ctr.Add)
		group.GET("/queryList", ctr.Query)
		group.DELETE("/del/:id", ctr.Del)
		group.PUT("/update", ctr.Update)
		group.POST("/addSong", ctr.SaveOrUpdatePlayListSongHandler)
		group.GET("/getSongs/:id", ctr.GetPlayListSongsHandler)

	}
}
