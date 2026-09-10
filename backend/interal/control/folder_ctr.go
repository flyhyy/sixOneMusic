package control

import (
	"log"

	"fnMusicServe/interal/config"
	"fnMusicServe/interal/model"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type FolderCtrImpl struct {
	FolderService service.FolderSceFace
}

func NewFolderCtr(srv service.FolderSceFace) *FolderCtrImpl {
	return &FolderCtrImpl{FolderService: srv}
}

// 查询操作
func (f *FolderCtrImpl) QueryHandler(g_ctx *gin.Context) {
	data, err := f.FolderService.Query()
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())

		return
	}
	utils.Success(g_ctx, data)
}

// 写入操作
func (f *FolderCtrImpl) WriteHandler(g_ctx *gin.Context) {
	var req []model.FolderWriteItemRequest

	err := g_ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	err = f.FolderService.Write(req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	utils.Success(g_ctx, nil)
}

// 删除操作
func (f *FolderCtrImpl) DelHandler(g_ctx *gin.Context) {
	var req model.FolderPathDeleteItemRequest

	err := g_ctx.ShouldBindQuery(&req)
	log.Printf("=======req=========%v", req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}

	err = f.FolderService.Delete(req)
	if err != nil {
		utils.Error(g_ctx, config.CodeError, err.Error())
		return
	}
	utils.Success(g_ctx, nil)
}
