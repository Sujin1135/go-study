package main

import (
	"fmt"
	"strconv"
)

func square(x int) int {
	return x * x
}

func fibonacci(n int) int {
	if n < 0 {
		return 0
	} else if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func AtoI(input string) (int, error) {
	if sv, err := strconv.Atoi(input); err != nil {
		return 0, nil
	} else {
		return sv, err
	}
}

func main() {
	num := 9
	fmt.Printf("%d x %d = %d\n", num, num, square(num))
}
