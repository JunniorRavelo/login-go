package handlers

import (
	"encoding/json"
	"net/http"
	"go-jwt-api/utils"
	"go-jwt-api/models"
)

// Login maneja la autenticación de usuarios decodificando el cuerpo de la solicitud en una estructura User,
// validando las credenciales y generando un token JWT si las credenciales son válidas.
// Si las credenciales son inválidas, devuelve una respuesta de error no autorizada.
//
// Parámetros:
//   - w: http.ResponseWriter para escribir la respuesta
//   - r: *http.Request que contiene la solicitud de inicio de sesión
//
// Respuestas:
//   - 200 OK: Devuelve un objeto JSON con el token JWT generado si las credenciales son válidas
//   - 400 Bad Request: Devuelve un error si el cuerpo de la solicitud es inválido
//   - 401 Unauthorized: Devuelve un error si las credenciales son inválidas
//   - 500 Internal Server Error: Devuelve un error si hay un problema al generar el token

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validar usuario (aquí solo una validación básica)
	if user.Username == "admin" && user.Password == "password" {
		token, err := utils.GenerateJWT(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		
		// Devolver el token si las credenciales son válidas
		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Devolver error si las credenciales son inválidas
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
