package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(
			w,
			"failed to encode response",
			http.StatusInternalServerError,
		)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	WriteJSON(w, statusCode, ErrorResponse{
		Error: message,
		Code:  statusCode,
	})
}
