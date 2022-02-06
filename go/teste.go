package main

import (
	"fmt"
	"sync"
)

var bola = make(chan int, 1)

func j1(wg *sync.WaitGroup, mutexGroup, end chan bool) {
	defer wg.Done()

	finished := false
	for !finished {
		select {
		case finished = <- end
		}
	}
}

func main() {
	wg := new(sync.WaitGroup)
	var mutex sync.Mutex
	bola <-2
	wg.Add(1)
	go j1(wg, &bola)
	wg.Wait()
	close(bola)
	fmt.Println(<-bola)
	
}
