package main

import "fmt"

func f() {
	fmt.Println("Start f() func")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic was recovered -", r)
		}
	}()

	g()
	fmt.Println("The end f() func")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/3 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다")
	}
	return a / b
}

// 복구는 안쓰는걸 추천
// Golang은 SEH를 지원하지 않는다
// 	- 성능
//	- 에러를 먹어 버리는 문제(오히려 에러 처리를 등한시)
func main() {
	fmt.Println("Start main func")
	f()
	fmt.Println("End main func")
}
