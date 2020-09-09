package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func HandlerRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
	// Permite utilizar un writer para pasar el mensaje requerido.
}

func HandlerHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is the API handler.")
	// Permite utilizar un writer para pasar el mensaje requerido.
}

func PostRequest(writer http.ResponseWriter, request *http.Request) {
	var metadata Metadata
	decoder := json.NewDecoder(request.Body)
	error := decoder.Decode(&metadata) // Recibe una interfaz para parsear la respuesta.

	if error != nil {
		fmt.Fprintf(writer, "[postRequest] %v", error)
		return
	}

	fmt.Fprintf(writer, "[payload] %v\n", metadata)
}

func UserPostRequest(writer http.ResponseWriter, request *http.Request) {
	var user User
	decoder := json.NewDecoder(request.Body)
	error := decoder.Decode(&user) // Recibe una interfaz para parsear la respuesta.
	response, err := user.ToJson()

	if error != nil {
		fmt.Fprintf(writer, "[postRequest] %v", error)
		return
	}

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(writer, "[payload] %v\n", user)
	writer.Write(response)
}