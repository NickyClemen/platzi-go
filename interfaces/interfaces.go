package main

import "fmt"

// Declaración de interface.
type animal interface {
	mover() string /* Para considerarse interface "animal", los structs tienen que tener un método
	con el mismo nombre. No hay que poner palabras clave. */
}

type perro struct {}
type pez struct {}
type pajaro struct {}

func (perro) mover() string {
	return "Soy un perro y camino."
}

func (pez) mover() string {
	return "Soy un pez, y estoy nadando."
}

func (pajaro) mover() string {
	return "Soy un pájaro, y estoy volando."
}

func moverAnimal(an animal) {
	fmt.Println(an.mover())
}

func main() {
	perro := perro {}
	pez := pez {}
	pajaro := pajaro {}

	moverAnimal(perro)
	moverAnimal(pez)
	moverAnimal(pajaro)
}