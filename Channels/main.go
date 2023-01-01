package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan time.Duration)
	go bucle(canal)
	fmt.Println("Llegue hasta aqui")

	msg := <-canal
	fmt.Println(msg)

}

func bucle(canal chan time.Duration) {
	inicio := time.Now()

	var i int
	for i = 0; i < 100000000000; i++ {

	}
	final := time.Now()
	canal <- final.Sub(inicio)
}
