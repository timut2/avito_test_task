package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/timut2/avito_test_task/internal/utils"
)

type ContextKey string

const UserIDKey ContextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/auth" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid authorization format")
			return
		}

		tokenString = parts[1]

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			if err == utils.ErrTokenExpired {
				utils.SendErrorResponse(w, http.StatusUnauthorized, "Token has expired")
			} else {
				utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			}
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
