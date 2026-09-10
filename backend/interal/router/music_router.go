package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitMusicRouter(apiGroup *gin.RouterGroup, ctr *control.MusicCtrImpl) {
	group := apiGroup.Group("/music")
	group.Use(middleware.JWTAuth())
	{
		group.GET("/all", ctr.QueryAllHandle)
		group.PUT("/collect", ctr.CollectHandler)
		group.GET("/collectSongs", ctr.GetUserSongsCollect)
	}
}
