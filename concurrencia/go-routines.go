package main

import (
	"net/http"
	"fmt"
	"time"
)

/* Palabra reservada:
	-go: Crea un nuevo hilo de ejecución que se ejecuta de manera paralela.package concurrencia
Hay un hilo principal, y otros paralelos.
Reduce el tiempo de ejecución.

Channels
No hay manera por defecto para que las distintas subrutinas de Go se comuniquen. Los channels permiten comunicar
subrutinas.
Da la oportunidad de enviar como recibir, por lo que es bidireccional.
Se envía con el operaodor "<-":
	Channel <- status */

func revisarServidor(server string, channel chan string) {
	_, error := http.Get(server)

	if error != nil {
		// fmt.Println(server, "no se encuentra disponible.")
		channel <- server + " no se encuentra disponible."
		// Transmisión de información al canal.
	}

	// fmt.Println(server, "está funcionando normalmente.")
	channel <- server + " está funcionando normalmente."
}

func iterarServidores(servers []string, channel chan string) {
	for _, server := range servers {
		go revisarServidor(server, channel)
	}
}

func main() {
	initExec := time.Now()
	// Retorna la hora actual.

	channel := make(chan string)
	// Declarar un channel.

	servidores := []string {
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	} /* Se ejecuta de manera secuencial. El programa no es responsable del tiempo de espera, por
	lo que queda en espera (queda bloqueado). */

	// iterarServidores(servidores, channel)

	/* Leer un canal. Algo tiene que estar pendiente del channel.
	Hay que especificar que esté pendiente del canal por todo el tiempo de ejecución del programa.
		fmt.Println(<- channel) */

	/* for i := 0; i < len(servidores); i++ {
		fmt.Println(<- channel)
	} // Para que esté pendiente de todas las subrutinas. */

	// Contador para romper el for infinito.
	i := 0

	for {
		if i > 10 {
			break
		}

		iterarServidores(servidores, channel)
		// "Dormir" el programa por un tiempo determinado.
		time.Sleep(1*time.Second)
		fmt.Println(<- channel)

		i++
	} /* No existe el ciclo while/do-while en Go. Para crear ciclos infinitos, hay que ajustar for.
	Para hacerlo indeterminado, se declara for {} (sin condiciones). */

	finishExec := time.Since(initExec) // Retorna diferencia horaria.
	fmt.Printf("Tiempo de ejecución transcurrido %s.", finishExec)
}