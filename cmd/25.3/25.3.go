package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	wg.Add(1)
	go sqaure(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

func sqaure(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Sqaure:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
