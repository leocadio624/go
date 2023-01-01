package main

import "fmt"

var numero int
var texto string
var status bool = true

func main() {
	fmt.Println("Hola mundo")
	numero2, numero3, numero4, texto, status := 2, 3, 4, "ejemplo texto", false

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
	showStatus()
}

func showStatus() {
	fmt.Println(status)
}
