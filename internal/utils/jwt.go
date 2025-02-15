package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretword = []byte("my_secret_key")

var (
	ErrTokenExpired = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretword)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	log.Println("Generated Token:", tokenString)
	log.Printf("Token Claims: UserID = %d, ExpiresAt = %v\n", claims.UserID, claims.ExpiresAt.Time)

	return tokenString, nil
}

func ValidateToken(tokenString string) (int, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretword, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, ErrTokenExpired
		}
		return 0, ErrInvalidToken
	}

	if !token.Valid {
		return 0, ErrInvalidToken
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return 0, ErrTokenExpired
	}

	return claims.UserID, nil
}
