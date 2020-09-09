package main

import (
	"net/http"
	"fmt"
	"time"
)

func revisarServidor(server string) {
	_, error := http.Get(server)

	if error != nil {
		fmt.Println(server, "no está disponible.")
	}

	fmt.Println(server, "está funcionando normalmente.")
}

func iterarServidores(servers []string) {
	for _, server := range servers {
		revisarServidor(server)
	}
}

func main() {
	initExec := time.Now()
	// Retorna la hora actual.

	servidores := []string {
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	} /* Se ejecuta de manera secuencial. El programa no es responsable del tiempo de espera, por
	lo que queda en espera (queda bloqueado). */

	iterarServidores(servidores)

	finishExec := time.Since(initExec) // Retorna diferencia horaria.
	fmt.Printf("Tiempo de ejecución transcurrido %s.", finishExec)
}