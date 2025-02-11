package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretword = []byte("my_secret_key")

func GenerateJwt(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecretword)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
