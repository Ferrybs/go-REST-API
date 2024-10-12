package middlewares

import (
	"blog/api/src/domain"
	"blog/api/src/infrastructure/auth"
	"blog/api/src/utils"
	"context"
	"net/http"
	"strings"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			var responseJson domain.ResponseJson = domain.ResponseJson{
				Ok:      false,
				Message: "Missing Authorization header",
			}
			utils.JsonWriter(w, http.StatusUnauthorized, responseJson)
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		if tokenString == "" {
			var responseJson domain.ResponseJson = domain.ResponseJson{
				Ok:      false,
				Message: "Missing token",
			}
			utils.JsonWriter(w, http.StatusUnauthorized, responseJson)
			return
		}
		userID, err := auth.ValidateToken(tokenString)
		if err != nil {
			var responseJson domain.ResponseJson = domain.ResponseJson{
				Ok:      false,
				Message: err.Error(),
			}
			utils.JsonWriter(w, http.StatusUnauthorized, responseJson)
			return
		}
		ctx := context.WithValue(r.Context(), domain.UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
