package middleware

import (
	"net/http"
	"strings"

	"github.com/Roman77St/salzo/internal/authctx"
	"github.com/Roman77St/salzo/internal/http/response"
	"github.com/Roman77St/salzo/internal/service/auth"
)

func Auth(
	authService *auth.Service,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			header := r.Header.Get("Authorization")

			if header == "" {
				response.WriteAppError(w, auth.ErrUnauthorized)
				return
			}

			const prefix = "Bearer "

			if !strings.HasPrefix(header, prefix) {
				response.WriteAppError(w, auth.ErrUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(
				header,
				prefix,
			)

			claims, err := authService.ParseToken(tokenString)
			if err != nil {
				response.WriteAppError(w, err)
				return
			}

			ctx := authctx.WithUser(
				r.Context(),
				authctx.User{
					ID:   claims.UserID,
					Role: claims.Role,
				},
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		})
	}
}
