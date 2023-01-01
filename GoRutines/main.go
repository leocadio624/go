package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//Indica que la funcion se invocara de forma asincrona
	go deletrea("Josue Eduardo Bernal Leocadio")

	fmt.Println("Ingrese una cadena de texto")
	var x string
	fmt.Scanln(&x)

}

/*
Autor : Eduardo B
Fecha creacion : 04/12/2022
Descripccion : Funcion asincrona
*/
func deletrea(cadena string) {
	letras := strings.Split(cadena, "")
	for _, i := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
}
