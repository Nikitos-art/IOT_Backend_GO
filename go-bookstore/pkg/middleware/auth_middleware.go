package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Nikitos-art/go-bookstore/pkg/utils"
)

type key string

const UserKey key = "user"


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		// Validate + extract claims
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		// 👉 Put user info into request context
		ctx := context.WithValue(r.Context(), UserKey, claims.UserID)

		// Continue request with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}