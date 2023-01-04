package main

import (
	"html/template"
	"os"
)

/*
Autor : Eduardo Bernal
Fecha : 04/01/23
Descp : Estructura que se pasa a los templates
*/
type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

/*
Autor : Eduardo Bernal
Fecha : 04/01/23
Descp : Incrementa el index de un slice
*/
var funcs = template.FuncMap{
	"increment": func(num int) int {
		return num + 1
	},
}

func main() {
	//t, err := template.New("greeting").Parse("Hola, mi nombre es: {{.Name}} y tengo {{.Age}} anios")

	//person := &Person{"Eduardo", 27}
	//loadTemplate("greetings.txt", person)

	//person := &Person{"Eduardo", 27, []string{"leer", "ver peliculas", "programar"}}
	//person := &Person{"Eduardo", 27, []string{}}
	//loadTemplate("greetings2.txt", person)

	person := &Person{"Eduardo", 27, []string{"<p>dede</p>", "leer", "ver peliculas", "programar"}}
	loadTemplate("index.html", person)

}

/*
Autor : Eduardo Bernal
Fecha : 04/01/23
Descp : Muestra el contenido de un template en consola
*/
func loadTemplate(fileName string, data interface{}) {

	t, err := template.New(fileName).Funcs(funcs).ParseFiles("templates/" + fileName)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
