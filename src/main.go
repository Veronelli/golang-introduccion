package main

import "fmt"

func normalFucntion(message string) {
	fmt.Printf("%s\n", message)

}

func sumar(a int32, b int32) int32 {
	return a + b

}

func dobleReturn(a int) (c, d int) {
	return a, a * 2

}

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

	x := 50
	y := 10
	var z int

	//Suma
	result := x + y
	fmt.Println(result)

	//Resta
	result = x - y
	fmt.Println(result)

	//Multiplicacion
	result = x * y
	fmt.Println(result)

	//Divicion
	result = x * y
	fmt.Println(result)

	//Modulo
	result = x % y
	fmt.Println(result)

	//Incremental
	x++
	fmt.Println(x)

	//Decremental
	y--
	fmt.Println(y)

	x = 10
	y = 4
	z = 4

	//Calcular area de un trapecio
	result = ((x + y) * z) / 2
	fmt.Println(result)

	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %d cursos\n", nombre, cursos)
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de variable
	fmt.Printf("Tipo de variable: %T\n", message)

	//Fuciones
	normalFucntion("Hola")
	fmt.Println(sumar(40, 20))

	value1, value2 := dobleReturn(4)
	fmt.Println(value1 + value2)

}
