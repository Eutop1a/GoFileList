package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte("ABAB")

type Claims struct {
	Id       uint   `json:"id" `
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username, password string) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(time.Hour * 24)
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
