package main

import (
	"fmt"
	"github.com/NickyClemen/calculator"
)

func main() {
	calc := calculator.Calc {}

	fmt.Println("Ingresá un valor:")
	entradaUno := calculator.ReadInput()

	fmt.Println("Ingresá otro valor:")
	entradaDos := calculator.ReadInput()

	fmt.Println("Ingresá una operación:")
	operador := calculator.ReadInput()

	fmt.Println("El resultado es", calc.Operate(entradaUno, entradaDos, operador))
}