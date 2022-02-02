package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// context.Background() <<- 기본 컨텍스트
	// WithCancel <<- cancel 기능을 추가하는 메서드
	ctx, cancel := context.WithCancel(context.Background())

	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)

	cancel()
	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
