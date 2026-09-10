package router

import (
	"fnMusicServe/interal/control"

	"github.com/gin-gonic/gin"
)

func AuthRouter(apiGroup *gin.RouterGroup, ctr *control.AuthCtrImpl) {
	group := apiGroup.Group("/auth")

	{
		// 登录
		group.POST("/login", ctr.LoginHandler)
		// 注册

		group.POST("/register", ctr.RegisterHandler)
	}
}
