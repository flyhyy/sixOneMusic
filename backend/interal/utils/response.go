package utils

import (
	"fnMusicServe/interal/config"

	"github.com/gin-gonic/gin"
)

func Success(g_ctx *gin.Context, data interface{}) {
	g_ctx.JSON(config.CodeSuccess, gin.H{
		"code": config.CodeSuccess,
		"msg":  config.MsgSuccess,
		"data": data,
	})
}

func Error(g_ctx *gin.Context, code int, err string) {
	g_ctx.JSON(config.CodeSuccess, gin.H{
		"code": code,
		"msg":  err,
		"data": nil,
	})
}
