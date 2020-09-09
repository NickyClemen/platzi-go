package main

import (
	"fmt"
)

func main() {
	server := NewServer(":3000")

	server.Handle("GET", "/", HandlerRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandlerHome, CheckAuth(), Logging()))
	// Parámetro: Handler y el middleware.

	fmt.Println("Servidor corriendo en el puerto 3000.")
	server.Listen()
}