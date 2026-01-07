package middleware

import (
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(userIdKey, authType string, userIdParseFunc func(string) (string, error)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || parts[0] != authType {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			userID, err := userIdParseFunc(parts[1])
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIdKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
