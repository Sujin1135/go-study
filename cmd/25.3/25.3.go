package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	chSize := 10
	ch := make(chan int, chSize)

	wg.Add(1)
	go sqaure(&wg, ch)

	for i := 0; i < chSize; i++ {
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
