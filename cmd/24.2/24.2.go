package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d 부터 %d 까지 합계는 %d 입니다\n", a, b, sum)
	wg.Done()
}

func main() {
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go SumAtoB(1, 100000000)
	}
	wg.Wait()
}
