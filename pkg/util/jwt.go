package util

import (
	"gin-template/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type customClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var jwtSecret = []byte(setting.App.JwtSecret)

func GenerateToken(username string, password string) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour)
	claims := customClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    setting.Common.ProjectName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
