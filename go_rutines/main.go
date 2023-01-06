package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"example.com/rutinas/data"
)

func main() {

	start := time.Now()
	wg := &sync.WaitGroup{} //agrega cada rutina para esperarla
	m := &sync.RWMutex{}    //para crear zona de exclusion mutua

	for i := 0; i < 10; i++ {
		wg.Add(1)
		//go showGoRutine(i, wg)
		go readBook(i, wg, m)
	}
	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("%dms\n", duration)

}

/*
Autor : Eduardo Bernal
Fecha : 05/01/23
Descp : Funcion que se ejecuta en una rutina de go
*/
/*
func showGoRutine(id int, wg *sync.WaitGroup) {

	delay := rand.Intn(500)
	fmt.Printf("Gorutine #%d with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
*/

/*
Autor : Eduardo Bernal
Fecha : 05/01/23
Descp : Cambia de estado item de lista
*/
func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {

	//data.FinishBook(id)
	data.FinishBook(id, m)
	delay := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()

}
