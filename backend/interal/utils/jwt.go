package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var MySecret = []byte("fnMusic")

type TokenInfo struct {
	UserID   uint   `json:"user_Id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(userId uint, userName string) (string, error) {
	data := TokenInfo{
		UserID:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			// token 有效期 7 天，过期后前端收到 401 会跳转登录页
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   "fnMusic",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	return token.SignedString(MySecret)
}

// 解析token

func ParseToken(tokenStr string) (*TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenInfo{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("签名算法错误")
		}
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenInfo)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的token")
}
