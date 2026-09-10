package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitAlbumRouter(apiGroup *gin.RouterGroup, ctr *control.AlbumCtrImpl) {
	apiGroup.Use(middleware.JWTAuth())

	group := apiGroup.Group("/album")
	{
		group.GET("/list", ctr.GetAlbumListHandle)
		group.GET("/song/:albumName", ctr.GetAlbumSongListHandle)
	}
}
