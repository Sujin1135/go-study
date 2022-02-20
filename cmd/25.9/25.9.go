package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	seconds := time.Duration(5)
	ctx, _ := context.WithTimeout(context.Background(), seconds*time.Second)
	wg.Add(4)

	publisher := NewPublisher(ctx)

	chrisCtx := context.WithValue(ctx, "name", "chris")
	chris, err := NewSubscriber(chrisCtx)
	if err != nil {
		fmt.Errorf("Occurred an error when try to create chris:%e", err)
	}

	susanCtx := context.WithValue(ctx, "name", "susan")
	susan, err := NewSubscriber(susanCtx)
	if err != nil {
		fmt.Errorf("Occurred an error when try to create susan:%e", err)
	}

	publisher.Subscribe(chris.subscribeCh)
	publisher.Subscribe(susan.subscribeCh)

	go func() {
		tick := time.Tick(time.Second * 1)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello Message")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}
