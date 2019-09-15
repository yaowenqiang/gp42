package main

import (
	"fmt"
	"time"
)

//type funcType func(int, int) int

type funcType func(time.Time)

func main() {
	testa(1, 2, 1.1)
	testc(1, 2, 1.1, 4, 5, 6, 7, 8)
	f := func() int { return 7 }
	fmt.Printf("f() = %+v\n", f())

	f1 := func(t time.Time) time.Time { return t }

	fmt.Printf("f1(time.Now()) = %+v\n", f1(time.Now()))

	var timer funcType = CurrentTime
	timer(time.Now())

	funcType(CurrentTime)(time.Now())
}

func CurrentTime(start time.Time) {
	fmt.Println(start)
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
