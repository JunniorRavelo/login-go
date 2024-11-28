package models

// User es una estructura que representa un usuario con un nombre de usuario, contraseña y correo electrónico.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
