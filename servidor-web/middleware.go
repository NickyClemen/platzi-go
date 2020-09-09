package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
)

func CheckAuth() Middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := true

			fmt.Println("Checking Authentication...")

			if flag {
				fn(writer, request)
			} else {
				return
			}
		} // El closure es necesario para no perder el contexto.
	}
} // Los middleware son handlers que reciben otro handler para evaluar cierta lógica.

func Logging() Middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(request.URL.Path, time.Since(start)) // A diferencia de fmt, muestra la hora exacta.
			}()
			// Permite ejecutar un bloque de código hasta el final.

			fn(writer, request)
		}
	}
}