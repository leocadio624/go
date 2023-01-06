package data

import (
	"fmt"
	"sync"
)

type Book struct {
	Id       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "El perfume", false},
	{3, "The world of Ice and Fire", false},
	{4, "Teoria de la noche", false},
	{5, "Blanca Olmedo", false},
	{6, "El principito", false},
	{7, "100 años de soledad", false},
	{8, "El alquimista", false},
	{9, "El libro del centenario", false},
	{10, "Blade runner", false},
}

/*
Autor : Eduardo Bernal
Fecha : 05/01/23
Descp : Lee zona critica en rutina con RLock
*/
func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock()
	for i, b := range books {
		if b.Id == id {
			index = i
			book = b
		}
	}
	m.RUnlock()

	return index, book
}

/*
Autor : Eduardo Bernal
Fecha : 05/01/23
Descp : Escribe en zona critica en rutina con Lock
*/
func FinishBook(id int, m *sync.RWMutex) {

	i, book := findBook(id, m)
	if i < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Finished Book: %s\n", book.Title)
}
