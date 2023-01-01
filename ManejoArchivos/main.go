package main

import "log"

func main() {

	/*
		archivo := "prueba.txt"
		f, err := os.Open(archivo)

		//Se hace para liberar el buffer en memoria
		defer f.Close()
		if err != nil {
			fmt.Println("Error al abrir el archivo")
			os.Exit(1)

		}
	*/
	Panic()
}

func Panic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic :c \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}

}
