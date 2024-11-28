package utils

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("mi_clave_secreta")

// Creamos una función para generar un token JWT con el nombre de usuario proporcionado
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Creamos una función para validar un token JWT y devolver el token si es válido
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	return token, err
}
