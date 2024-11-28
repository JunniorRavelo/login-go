package main

import (
	"fmt"
	"log"
	"net/http"
	"go-jwt-api/handlers"
	"go-jwt-api/middleware"
	"github.com/gorilla/mux"
)

func main() {
	// Creamos un nuevo enrutador usando mux
	r := mux.NewRouter()

	// Definimos la ruta para el login, que maneja solicitudes POST
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	
	// Definimos la ruta para obtener el usuario, que maneja solicitudes GET y usa middleware para verificar JWT
	r.HandleFunc("/user", middleware.CheckJWT(handlers.GetUser)).Methods("GET")

	// Imprimimos un mensaje indicando que el servidor está corriendo en el puerto 8080
	fmt.Println("Servidor corriendo en el puerto 8080")
	
	// Iniciamos el servidor en el puerto 8080 y registramos cualquier error que ocurra
	log.Fatal(http.ListenAndServe(":8080", r))
}