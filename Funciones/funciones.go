package main

import "fmt"

func main() {

	/*
		fmt.Println(uno2(555))
		numero, estado := dos(1)

		fmt.Println(numero)
		fmt.Println(estado)
	*/
	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(5, 7, 6, 15))
	fmt.Println(calculo(1, 2, 3, 4, 5))

}

func uno(numero int) int {
	return numero * 2
}

func uno2(numero int) (z int) {
	z = numero * 2
	return

}

func dos(numero int) (int, bool) {

	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func calculo(parametros ...int) int {
	total := 0

	for _, val := range parametros {
		total += val

	}
	return total

}
