package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//firstAleatory()
	secondAleatory()

}

/*
Autor:Eduardo Bernal
Fecha:01/01/23
Descp:Genera numeros aleatorios(siempre los mismos)
*/
func firstAleatory() {
	var ale int

	fmt.Print("\n")

	for i := 0; i < 10; i++ {
		//Genera aleatorios entre 0 y 100
		ale = rand.Intn(101)
		fmt.Print(ale, "-")
	}
}

/*
Autor:Eduardo Bernal
Fecha:01/01/23
Descp:Genera numeros aleatorios usando la estructura NewSource
que se le pasa un parametro hora para generar aleatorios distintos
en cada ejecucion
*/
func secondAleatory() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ale int

	for i := 0; i < 10; i++ {
		ale = r.Intn(101)
		fmt.Print(ale, "-")

	}
	fmt.Println(" ")
}
