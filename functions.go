package main

import (
	"fmt"
)

func main() {
	testa(1, 2, 1.1)
	testc(1, 2, 1.1, 4, 5, 6, 7, 8)
}

func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return 1
		} else {
			return i
		}
	}

	return 1
}

func FunctionName(a int32, b float32) (t1 int32, t2 float32) {
	return a, b
}

func testa(a, b int, z float32) bool {
	i := 1
	fmt.Printf("i = %+v\n", i)
	return true
}

func testb(a, b int, z float32) bool {
	return true
}

func testc(a, b int, z float32, values ...int) bool {
	//可变参数
	fmt.Printf("values	 = %+v\n", values)
	return true
}
