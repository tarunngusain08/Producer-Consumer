package main

import (
	"fmt"
	"sync"
)

func Producer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		channel <- i
	}
}

func Consumer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val := <-channel:
			fmt.Println(val)
		default:
			break
		}
	}
}

func main() {
	wg := new(sync.WaitGroup)
	channel := make(chan int)
	defer close(channel)
    
	wg.Add(1)
	go Producer(channel, wg)
	wg.Add(1)
	go Consumer(channel, wg)

	wg.Wait()
}
