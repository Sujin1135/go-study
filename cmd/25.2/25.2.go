package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go square(ch)
	ch <- 9
	fmt.Println("Never print")
}

func square(ch chan int) {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")

		fmt.Println("value is", <-ch)
	}
}
