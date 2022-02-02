package main

import (
	"context"
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory\n")

	wg.Add(3)

	go InitBody(ctx, tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	time.Sleep(10 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("All Done")
}

func InitBody(ctx context.Context, ch chan *Car) {
	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
			ch <- &Car{Body: "Tesla Model 3"}
		case <-ctx.Done():
			close(ch)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)

		car.Tire = "Snow tire"
		paintCh <- car
	}

	close(paintCh)
	wg.Done()
}

func PaintCar(ch chan *Car) {
	for car := range ch {
		car.Color = "Yellow"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
