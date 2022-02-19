package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 2)

	wg.Add(2)

	go square(&wg, ch)
	go square(&wg, ch)

	ch <- 9
	ch <- 8

	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)

	fmt.Println("Square:", n*n)
	wg.Done()
}
