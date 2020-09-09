package main

import "fmt"

type perro struct {}
type pez struct {}
type pajaro struct {}

func (perro) caminar() string {
	return "Soy un perro y camino."
}

func (pez) nada() string {
	return "Soy un pez, y estoy nadando."
}

func (pajaro) vuela() string {
	return "Soy un pájaro, y estoy volando."
}

func moverPerro(dog perro) {
	fmt.Println(dog.caminar())
}

func moverPez(fish pez) {
	fmt.Println(fish.nada())
}

func moverPajaro(bird pajaro) {
	fmt.Println(bird.vuela())
}

func main() {
	perro := perro {}
	pez := pez {}
	pajaro := pajaro {}

	moverPerro(perro)
	moverPez(pez)
	moverPajaro(pajaro)
}