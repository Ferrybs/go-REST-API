package utils

import (
	"blog/api/src/domain"
	"encoding/json"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, statusCode int, responseJson domain.ResponseJson) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseMap := map[string]interface{}{
		"ok":      responseJson.Ok,
		"message": responseJson.Message,
	}
	if responseJson.Data != nil {
		responseMap["data"] = responseJson.Data
	}
	json.NewEncoder(w).Encode(responseMap)
}