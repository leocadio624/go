package main

import "fmt"

func main() {

	/*
		forma 1
			i := 1
			for i < 10 {
				fmt.Println(i)
				i += 1
			}
	*/

	/*
		forma 2
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	*/

	/*
		while
		numero := 0
		for {

			fmt.Println("Continuo")
			fmt.Println("Ingrese un numero")
			fmt.Scanln(&numero)

			if numero == 100 {
				break
			}

		}
	*/

	/*
		var i = 0
		for i < 10 {
			fmt.Printf("Valor de i: %d\n", i)

			if i == 5 {
				fmt.Printf("Multiplicamos por 2 \n")
				i = i * 2
				continue
			}

			fmt.Printf("Paso por aqui\n")
			i += 1
		}
	*/
	var i int = 0
rutina:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a rutina sumando 2 a i")
			goto rutina

		}
		fmt.Printf("Valor de i %d\n", i)
		i++
	}

}
