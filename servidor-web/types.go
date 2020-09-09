package main

import (
	"net/http"
	"encoding/json"
)

// Declaración de un tipo de dato custom.
type Middleware func(http.HandlerFunc) http.HandlerFunc
type Metadata interface {}
type User struct {
	Name string `json:"name"` // Indica que cuando sea decodificado como JSON, tiene que llevar ese alias.
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (user *User) ToJson() ([]byte, error) {
	return json.Marshal(user) // Permite encodear el struct en un JSON.
} // Parsear el struct a JSON.