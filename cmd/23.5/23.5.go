package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b must not be zero")
	}

	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(5, 2)
	divide(3, 0)
}
