package utils

import (
	"blog/api/src/domain"
	"errors"
	"net/http"
)

func GetUserIDFromContext(r *http.Request) (domain.UserID, error) {
	userID := r.Context().Value(domain.UserIDKey)
	if userID == nil {
		return "", errors.New("user ID not found in context")
	}
	if id, ok := userID.(domain.UserID); ok {
		return id, nil
	}
	return "", errors.New("invalid user ID type in context")
}
