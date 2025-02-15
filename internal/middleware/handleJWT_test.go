package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTMiddleware_MissingAuthorization(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
	rr := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Next handler should not be called")
	})

	middleware := JWTMiddleware(nextHandler)
	middleware.ServeHTTP(rr, req)

	// Проверяем, что статус ответа 401
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Authorization header is missing")
}

func TestJWTMiddleware_InvalidFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
	req.Header.Set("Authorization", "InvalidFormat")

	rr := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Next handler should not be called")
	})

	middleware := JWTMiddleware(nextHandler)
	middleware.ServeHTTP(rr, req)

	// Проверяем, что статус ответа 401
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid authorization format")
}
