package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetToken(username string) string {

	c := MyClaims{
		username, "AccessToken", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			Issuer:    Author,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, _ := token.SignedString((Secret))

	return res
}
