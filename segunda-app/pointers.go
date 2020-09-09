package main

import "fmt"

func main() {
	x := 25
	y := &x

	fmt.Println(x)
	cambiarValor(x)
	fmt.Println(&x) /* &: Retorna un valor hexagecimal. Es la representación del espacio de memoria
	dónde está almacenada la variable. */
	fmt.Println(y)
	fmt.Println(*y) /* *: Retorna el valor de ese espacio de memoria. */
}

func cambiarValor(a int) {
	fmt.Println(&a) // Se hace una copia del valor del parámetro.
	a = 36
}