package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	token, err := GenerateJWT(1)
	assert.NoError(t, err)

	userID, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, 1, userID)
}

func TestValidateToken_Expired(t *testing.T) {
	expirationTime := time.Now().Add(-1 * time.Hour)
	claims := &Claims{
		UserID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretword)
	assert.NoError(t, err)

	_, err = ValidateToken(tokenString)
	assert.Error(t, err)
	assert.Equal(t, ErrTokenExpired, err)
}
