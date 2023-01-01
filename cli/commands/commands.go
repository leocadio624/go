package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/cli/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Obtiene la cadena de entrada de la linea de comandos
*/
func GetInput() (string, error) {
	fmt.Print("-> ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Muestra cadena de resultados en linea de
comandos
*/
func ShowInConsole(expensesList []float32) {

	fmt.Println(contentString(expensesList))

}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Construye una cadena en buffer y la retorna
*/
func contentString(expensesList []float32) string {
	//Guarda una cadena en buffer
	builder := strings.Builder{}
	max, min, avg, total := expensesDetails(expensesList)

	fmt.Println("")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))

		if i == len(expensesList)-1 {
			builder.WriteString("")
			builder.WriteString("==================================\n")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
			builder.WriteString("==================================\n")
		}
	}

	return builder.String()
}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Realiza el calculo de las operaciones y retorna los valores
max, min, avg, total
*/
func expensesDetails(expensesList []float32) (max, min, avg, total float32) {
	if len(expensesList) == 0 {
		return
	}
	max = expenses.Max(expensesList...)
	min = expenses.Min(expensesList...)
	avg = expenses.Average(expensesList...)
	total = expenses.Sum(expensesList...)

	return
}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Exporta el resultado en un archivo txt usando una bandera en la funcion principal
*/
func Export(fileName string, list []float32) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()
	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil {
		return err
	}

	return w.Flush()
}
