package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* structs: Tipo de dato custom. Se le pueden bindear métodos. */
type calc struct {} // Definición de structs.

func (calc) operate(inputUno string, inputDos string, operator string) int {
	operadorUno := parseInt(inputUno)
	operadorDos := parseInt(inputDos)

	switch(operator) {
		case "+":
			return operadorUno + operadorDos
		case "-":
			return operadorUno - operadorDos
		case "*":
			return operadorUno * operadorDos
		case "/":
			return operadorUno / operadorDos
		default:
			return 0
	}
} // Receiver function. Bindea la función al struct.

func parseInt(input string) int {
	// operadorUno, _ := strconv.Atoi(valores[0])
	operador, _ := strconv.Atoi(input) /* Convierte un string a un entero. Devuelve dos elementos: El valor convertido en
	entero, y un error. Cuando hubo un error, lo devuelve, sino, retorna null (nil).
	El _ crea un wildcard. Permite omitir uno de los valores de retorno de una función. Si se usa otro valor,
	no compila el programa. */
	return operador
}

func readInput() string {
	// Scanner permite leer el input del usuario. Se lo declara, y se pasa el input de entrada de consola.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Scanea la consola.
	return scanner.Text() // Asigna un string.
}

func main() {
	fmt.Println("Ingresa un valor: ")
	inputUno := readInput()

	fmt.Println("Ingresa otro valor: ")
	inputDos := readInput()

	fmt.Println("Ingresa una operación: ")
	operador := readInput()

	calc := calc{}
	result := calc.operate(inputUno, inputDos, operador)

	fmt.Println("El resultado es", result)
}