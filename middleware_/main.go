package main

import "fmt"

var result int

func main() {
	fmt.Println("Inicio")

	result = opMiddleware(sumar)(2, 3)
	fmt.Println(result)
	result = opMiddleware(resta)(2, 3)
	fmt.Println(result)
	result = opMiddleware(multitiplicar)(2, 3)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}
func resta(a, b int) int {
	return a - b
}
func multitiplicar(a, b int) int {
	return a * b
}

func opMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)

	}
}
