package middleware

import (
	"fnMusicServe/interal/config"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

func AudioAuth() gin.HandlerFunc {
	return func(g_ctx *gin.Context) {
		tokenStr := g_ctx.Query("token")

		if tokenStr == "" {

			utils.Error(g_ctx, config.CodeAuth, "token 缺失")
			g_ctx.Abort()
			return

		}

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			utils.Error(g_ctx, config.CodeAuth, "无效或过期的token")
			g_ctx.Abort()
			return
		}

		g_ctx.Set("userID", claims.UserID)
		g_ctx.Set("userName", claims.UserName)

		g_ctx.Next()
	}
}
