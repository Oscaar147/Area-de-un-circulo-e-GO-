package main

import "fmt"

func main() {
	const pi float64 = 3.1416
	var radio float64 = 0
	var area float64 = 0

	fmt.Println("|---------------5------------------------------------------------|")
	fmt.Println("| Hola Bienvenido al programa para sacar area de un circulo!!!   |")
	fmt.Println("|----------------------------------------------------------------|")
	fmt.Println("Ingresa la medida del radio del circulo: ")
	fmt.Scan(&radio)

	area = pi * (radio * radio)

	fmt.Println("El area del triangulo es: ", area)

}
