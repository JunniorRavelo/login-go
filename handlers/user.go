package handlers

import (
	"encoding/json"
	"net/http"
	"go-jwt-api/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Simular un usuario obtenido de base de datos
	user := models.User{
		Username: "admin",
		Email:    "admin@example.com",
	}

	// Devolver el usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
