package middleware

import (
	"context"
	"github.com/aririfani/auth-app/config"
	"github.com/aririfani/auth-app/internal/pkg/chicustom"
	"github.com/aririfani/auth-app/internal/pkg/token"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func JWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		cfg := config.NewConfig()
		if !strings.Contains(authorizationHeader, "Bearer") {
			chicustom.WriteError(w, r, jwt.NewValidationError("Invalid Authorization Header", jwt.ValidationErrorUnverifiable))
			return
		}

		tokenStr := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims, err := token.New(token.WithSecretKey(cfg.GetString("app.secret_key"))).GetClaims(tokenStr)
		if err != nil {
			chicustom.WriteError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
