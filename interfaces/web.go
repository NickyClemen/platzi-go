package main

import (
	"net/http"
	"fmt"
	"io"
)

type webWriter struct {}

func (webWriter) Write(param []byte) (int, error) {
	fmt.Println(string(param))
	return len(param), nil
	// len(param): Longitud de bytes.
} // Por parámetro, tiene un slice de tipo byte.

func main() {
	response, error := http.Get("http://google.com")
	/* Retorna un response y un error. */
	writer := webWriter {}

	if error != nil {
		fmt.Println(error)
	}

	/* Copy (método nativo). Cuando se crea un request, lo que tiene response.Body, es un Reader (no tiene la respuesta,
	sino una manera de ir leyéndola. Permite leer el flujo de bytes, y parsear la respuesta).
	Writer: Interface. Hay que crear un struct que implemente el méodo "write" para poder utilizarla.
	Los métodos son muy genéricos, por lo que hay que implementar las interfaces. */
    io.Copy(writer, response.Body)
}