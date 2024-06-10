package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/amnestia/tnderlike/internal/lib/paseto"
	"github.com/amnestia/tnderlike/pkg/response"
)

// AuthorizationModule auth module
type AuthorizationModule struct {
	Token paseto.Handler
}

// Auth authorization token check
func (a *AuthorizationModule) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", "Invalid token").WriteJSON(w)
			return
		}
		token := strings.Replace(authHeader, "Bearer ", "", -1)
		payload, err := a.Token.Extract(token)
		if err != nil {
			response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", err.Error()).WriteJSON(w)
			return
		}
		if payload.TokenType != paseto.AccessToken {
			response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", "Invalid token").WriteJSON(w)
			return
		}

		ctx = context.WithValue(ctx, paseto.AuthData, payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
