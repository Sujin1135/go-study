package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Color string
	Tire  string
}

var wg sync.WaitGroup
var startTime = time.Now()
var limit = 10

func main() {
	wg.Add(3)
	tireCh := make(chan *Car, limit)
	paintCh := make(chan *Car, limit)

	go makeBody(tireCh, time.Duration(limit), &wg)
	go installTire(tireCh, paintCh, &wg)
	go paintCar(paintCh, &wg)

	wg.Wait()
}

func makeBody(tireCh chan *Car, endSeconds time.Duration, wg *sync.WaitGroup) {
	tick := time.Tick(time.Second)
	terminate := time.After(endSeconds * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{Body: "Make the body"}
			tireCh <- car
		case <-terminate:
			fmt.Println("Work done")
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func installTire(tireCh chan *Car, paintCh chan *Car, wg *sync.WaitGroup) {
	for c := range tireCh {
		time.Sleep(time.Second)
		c.Tire = "Snow Tire"
		paintCh <- c
	}

	close(paintCh)
	wg.Done()
}

func paintCar(paintCh chan *Car, wg *sync.WaitGroup) {
	for c := range paintCh {
		time.Sleep(time.Second)
		c.Color = "Blue"
		fmt.Println("The car aws made", c)
	}

	wg.Done()
}
