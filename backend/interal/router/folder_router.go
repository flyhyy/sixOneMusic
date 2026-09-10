package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitFolderRouter(apiGroup *gin.RouterGroup, ctr *control.FolderCtrImpl) {
	group := apiGroup.Group("/folder")
	group.Use(middleware.JWTAuth())
	{

		group.GET("/query", ctr.QueryHandler)
		group.POST("/write", ctr.WriteHandler)
		group.DELETE("/del", ctr.DelHandler)
	}
}
