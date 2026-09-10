package control

import (
	"fnMusicServe/interal/config"
	"fnMusicServe/interal/service"
	"fnMusicServe/interal/utils"

	"github.com/gin-gonic/gin"
)

type ScanCtrImpl struct {
	ScanSer   service.ScanSerFace
	FolderSer service.FolderSceFace
}

func NewScanCtr(scan_ser service.ScanSerFace, folder_ser service.FolderSceFace) *ScanCtrImpl {
	return &ScanCtrImpl{ScanSer: scan_ser, FolderSer: folder_ser}
}

func (s *ScanCtrImpl) ScanHandler(g_ctx *gin.Context) {
	arr, err := s.FolderSer.Query()
	if err != nil {
		utils.Error(g_ctx, config.CodeError, "查询路径出错")
	}
	if len(arr) == 0 {
		utils.Success(g_ctx, "未设置文件路径")
	} else {
		var paths []string

		for _, item := range arr {
			paths = append(paths, item.Path)
		}

		s.ScanSer.Scan(paths)
		utils.Success(g_ctx, "开始扫描")

	}
}

func (s *ScanCtrImpl) ScanStatus(g_ctx *gin.Context) {
	utils.Success(g_ctx, s.ScanSer.GetScanStatus())
}
