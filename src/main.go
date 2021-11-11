package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.24
	const pi2 = 3.54

	fmt.Println("pi ", pi)
	fmt.Println("pi2 ", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Calcular un area de un cuadrado
	const baseCuadrado int = 10
	const alturaCuadrado int = 20
	fmt.Println(baseCuadrado * alturaCuadrado)
}
