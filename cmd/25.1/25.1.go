package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go sqaure(&wg, ch)
	ch <- 9
	wg.Wait()
}

func sqaure(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Sqaure:", n*n)
	wg.Done()
}
