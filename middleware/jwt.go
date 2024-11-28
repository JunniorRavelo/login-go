package middleware

import (
	"net/http"
	"strings"
	"go-jwt-api/utils"
)

// CheckJWT es un middleware que verifica si un token JWT es válido antes de permitir el acceso a un manejador.
func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// Verificar si el encabezado de autorización está presente
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Verificar si el token es válido
		tokenString := strings.Split(authHeader, " ")[1]
		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Pasar al siguiente manejador
		next(w, r)
	}
}
