package router

import (
	"fnMusicServe/interal/control"

	"github.com/gin-gonic/gin"
)

func InitAudioRouter(apiGroup *gin.RouterGroup, ctr *control.AudioCtrImpl) {
	// apiGroup.Use(middleware.AudioAuth())
	group := apiGroup.Group("/audio")
	{
		group.GET("/play/:id", ctr.AudioPlayHandler)
		group.GET("/play/lrc/:id", ctr.AudioLrcHandler)
	}
}
