package handler

import (
	"net/http"

	"github.com/Roman77St/selzo/internal/service/auth"

)

type AuthHandler struct {
	service auth.Service
}

func NewAuthHandler(service *auth.Service) *AuthHandler {
	return &AuthHandler{
		service: *service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	http.Error(
		w,
		http.StatusText(http.StatusNotImplemented),
		http.StatusNotImplemented,
	)
}
