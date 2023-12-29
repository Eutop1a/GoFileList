package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username  string `json:"username"`
	TokenType string `json:"token-type"`
	jwt.StandardClaims
}
type Token struct {
	TokenString string // token
}

type Result struct {
	UserName string // 返回的username
	Status   string // 状态：
	// case 0: "" (success)
	// case 1: "Token has expired"
	// case 2: "Error parsing token:"
}

const TokenTime = time.Hour * 10

var Secret = []byte("Author:Eutop1a")

const Author = "Eutop1a"
