package utils

import (
	"errors"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 密码加密
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// 密码比对
func PasswordDiff(hashPassword string, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword))
}

// // 匹配 音频格式
// func RegAudio(name string) bool {
// 	return RegxAudio.MatchString(name)
// }

// // 匹配歌词格式
// func RegLrc(name string) bool {
// 	return RegxLrc.MatchString(name)
// }

// 匹配 音频格式
func MatchAudio(name string) bool {
	audioExts := []string{
		".mp3",
		".aac",
		".wav",
		".flac",
		".ogg",
		".m4a",
		".wma",
		".alac",
		".aiff",
		".opus",
		".amr",
		".ac3",
	}
	ext := strings.ToLower(filepath.Ext(name))

	return slices.Contains(audioExts, ext)
}

// 匹配歌词格式
func MatchLrc(name string) bool {
	lrcExts := []string{".lrc"}

	ext := strings.ToLower(filepath.Ext(name))

	return slices.Contains(lrcExts, ext)
}

// 分页
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {

		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10

		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// 获取userID
func GetUserID(g_ctx *gin.Context) (uint, error) {
	id, exist := g_ctx.Get("userID")

	if !exist {
		return 0, errors.New("无法正确获取userID")
	}
	userID, ok := id.(uint)
	if !ok {
		return 0, errors.New("userID 类型不正确")
	}

	return userID, nil
}

// string转换为 uint
func StrToUnit(str string) (uint, error) {
	if str == "" {
		return 0, errors.New("字符串为空")
	}

	uintVal, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}

	return uint(uintVal), nil
}
