package main

import ("fmt")

func main() {
	/* Un Map es una estructura key: value. */
	mapOne := make(map[string]int) // Inicialización del map. [tipo_de_dato_key]tipo_de_dato_value.
	mapOne["a"] = 8

	fmt.Println(mapOne["a"])
}