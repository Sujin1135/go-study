package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	num := 9
	fmt.Printf("%d x %d = %d\n", num, num, square(num))
}
