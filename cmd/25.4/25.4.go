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

	for i := 0; i < 10; i++ {
		ch <- i * i
	}

	wg.Wait()
}

func sqaure(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // 일정 시간마다 신호를 주는 채널
	terminate := time.After(10 * time.Second) // 일정 시간 후 한번 신호를 주는 채널

	for {
		// animation 60fps + input
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square:", n*n)
			time.Sleep(time.Second)
		}
	}
}
