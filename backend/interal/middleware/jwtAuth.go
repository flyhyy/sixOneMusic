package middleware

import (
	"log"
	"strings"

	"fnMusicServe/interal/config"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(g_ctx *gin.Context) {
		// 获取 Authorization 请求头
		ahthHeader := g_ctx.GetHeader("Authorization")
		log.Printf("======ahthHeader========%v", ahthHeader)
		if ahthHeader == "" {
			utils.Error(g_ctx, config.CodeAuth, "token缺失")
			g_ctx.Abort()
			return

		}
		// 分割

		strs := strings.SplitN(ahthHeader, " ", 2)
		log.Printf("======strs========%v", len(strs))

		if !(len(strs) == 2 && strs[0] == "Bearer") {

			utils.Error(g_ctx, config.CodeAuth, "Authorization 格式错误，应为 Bearer <Token>")
			g_ctx.Abort()
			return
		}
		tokenStr := strs[1]
		log.Printf("======tokenStr========%v", tokenStr)

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
