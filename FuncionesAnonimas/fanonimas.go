package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma de 5 + 7 = %d\n", Calculo(5, 7))

	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Suma de 5 - 7 = %d\n", Calculo(5, 7))

	Operaciones()

}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())

}
