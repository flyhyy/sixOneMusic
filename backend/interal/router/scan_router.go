package router

import (
	"fnMusicServe/interal/control"
	"fnMusicServe/interal/middleware"

	"github.com/gin-gonic/gin"
)

func InitScanRouter(apiGroup *gin.RouterGroup, ctr *control.ScanCtrImpl) {
	group := apiGroup.Group("/scan")
	group.Use(middleware.JWTAuth())

	{
		group.GET("/handler", ctr.ScanHandler)
		group.GET("/status", ctr.ScanStatus)
	}
}

// 2rs6 vmbg gnj4 lnbr 7eii 7ckg qclf hvib
// LDXP-74PR4LBHVWNQ
