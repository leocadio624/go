package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//readFile()
	//readFile2()
	//createFile()
	grabarArchivo()
}

/*
Autor : Eduardo B
Fecha creacion : 01/12/2022
Descripccion : Lee el contenido de un archivo
*/
func readFile() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}

}

/*
Autor : Eduardo B
Fecha creacion : 01/12/2022
Descripccion : Lee el contenido de archivo con Scanner
*/
func readFile2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea>\n" + registro + "\n")
		}

	}
	archivo.Close()
}

/*
Autor : Eduardo B
Fecha creacion : 01/12/2022
Descripccion : Crea|sobreescribe un archivo
*/
func createFile() {
	archivo, err := os.Create("./NuevoArchivo")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una nueva linea")
	archivo.Close()

}

/*
Autor : Eduardo B
Fecha creacion : 01/12/2022
Descripccion : Invoca funcion Append()
*/
func grabarArchivo() {
	fileName := "./NuevoArchivo"
	if Append(fileName, "\nEsta es una segunda linea") == false {
		fmt.Println("Error al grabar la segunda linea")
	}
}

/*
Autor:Eduardo B
Fecha creacion:01/12/2022
Descripccion : Agrega una nueva linea al archivo sin
borrar el contenido del archivo
*/
func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Hubo un error al abrir el archivo")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error al agregar una nueva linea")
		return false
	}
	return true

}
