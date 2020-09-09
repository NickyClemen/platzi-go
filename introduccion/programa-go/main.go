package main /* Forma de almacenar código. Se puede poner cualquier nombre. */

import "fmt" /* Impotar librerías necesarias para la ejecución del programa. */

func main() {
	var mensaje string = "Hello world!" // Declaración explícita de variable: var [varName] [tipo de dato].
	mensajeFacil := "Hello world! :=" // Forma implícita de asignar tipo de dato. Es la más popular.

	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)
	/* Tipos de dato:
		1) string
		2) float64
		3) int
		4) boolean */

	a := 10.
	var b float64 = 3
	fmt.Println((a / b))

	var c int = 10
	d := 3
	fmt.Println(c / d)

	x := true
	y := false

	fmt.Println((x || y)) // AND.
	fmt.Println(y && x) // OR.
	fmt.Println(!x) // Negar.

	privada()
	Publica()
	imprimirHola()
} // Punto de inicio del programa. Le indica al compilador qué tiene que ejecutar.


// Diferencia si una función es privada o pública por la primera letra: Si es mayúscula, es pública, sino, es privada.
func privada() {
	fmt.Println("Ejecutar lógica que no necesita ser exportada.");
}

func Publica() {
	fmt.Println("Lógica que quiero exportar.");
}

func imprimirHola() {
	defer fmt.Println("Hello") // Fragmento de código que no debe ser ejecutado hasta que no termine la ejecución de la función.
	fmt.Println("World")
}