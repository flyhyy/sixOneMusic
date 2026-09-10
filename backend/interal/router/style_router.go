package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitStyleRouter(apiGroup *gin.RouterGroup, ctr *control.StyleCtrImpl) {
	apiGroup.Use(middleware.JWTAuth())

	group := apiGroup.Group("/style")

	{
		group.GET("/list", ctr.StyleQueryHandle)
		group.GET("/song/:styleId", ctr.StyleQuerySongHandle)
	}
}
