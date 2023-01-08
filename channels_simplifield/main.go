package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {

	wg := &sync.WaitGroup{}
	ids_channel := make(chan string)
	fake_channel := make(chan string)
	closed_channel := make(chan int)

	wg.Add(3)

	go generateIDS(wg, ids_channel, closed_channel)
	go generateFakeIDS(wg, fake_channel, closed_channel)

	go logIds(wg, ids_channel, fake_channel, closed_channel)

	wg.Wait()

}

func generateFakeIDS(wg *sync.WaitGroup, fake_channel chan string, closed_chanel chan int) {
	for i := 0; i < 100; i++ {
		id := uuid.New()
		fake_channel <- fmt.Sprintf("%d. %s", i+1, id.String())

	}
	close(fake_channel)
	closed_chanel <- 1
	wg.Done()
}

func generateIDS(wg *sync.WaitGroup, ids_chan chan string, closed_chanel chan int) {

	//id := uuid.New()
	//ids_chan <- id.String()

	for i := 0; i < 100; i++ {
		id := uuid.New()
		ids_chan <- fmt.Sprintf("%d. %s", i+1, id.String())

	}
	close(ids_chan)
	closed_chanel <- 1
	wg.Done()

}

func logIds(wg *sync.WaitGroup, ids_chan <-chan string, fake_ids_chan <-chan string, closed_chan chan int) {

	closedCounter := 0
	for {
		select {
		case id, ok := <-ids_chan:
			if ok {
				fmt.Println("ID:", id)
			}
		case id, ok := <-fake_ids_chan:
			if ok {
				fmt.Println("Fake ID:", id)
			}
		case count, ok := <-closed_chan:
			if ok {
				closedCounter += count

			}
		}

		if closedCounter == 2 {
			close(closed_chan)
			break
		}
	}

	/*
		for id := range ids_chan {
			fmt.Println(id)
		}
	*/
	wg.Done()

}
