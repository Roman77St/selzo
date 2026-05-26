package handler

import (
	"net/http"
)

// HealthHandler returns a simple "OK" response to indicate the service is healthy.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}