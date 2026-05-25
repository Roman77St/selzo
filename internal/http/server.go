package http

import (
	"net/http"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()

	return &http.Server{
		Addr:    addr,
		Handler: mux,
		ReadTimeout:       5, // 5 seconds
		WriteTimeout:      10, // 10 seconds
		IdleTimeout:       2 * 60, // 2 minutes
	}
}