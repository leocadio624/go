package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		paises := make(map[string]string)
		paises["mexico"] = "DF"
		paises["argentina"] = "Buenos aires"

		for key, element := range paises {
			fmt.Println("Key:", key, "=>", "Element:", element)

		}
	*/

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 39}

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 55

	delete(campeonato, "Boca Juniors")

	for key, element := range campeonato {
		fmt.Printf("El equipo %s tiene un putaje %d\n", key, element)

	}

	//Pregunta si existe un equipo por nombre
	puntaje, existe := campeonato["Mineros"]

	fmt.Println(puntaje)
	fmt.Println(existe)

	fmt.Println(reflect.TypeOf(existe))
	fmt.Println(reflect.TypeOf(puntaje))

}
