package middleware

import (
	"github.com/aririfani/auth-app/fetch-service/internal/pkg/chicustom"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AdminRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("claims").(jwt.MapClaims)["Role"].(string)
		if userRole != "admin" {
			chicustom.WriteError(w, r, jwt.NewValidationError("role is not allowed", 403))
			return
		}

		next.ServeHTTP(w, r)
	})
}
