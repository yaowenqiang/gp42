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
	list(1, 2, 3, 45, 6, 7, 78)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list(numbers...)

	fplus := func(x, y int) int { return x + y }
	fmt.Printf("fplus(3,4) = %+v\n", fplus(3, 4))

	a := func(x, y int) int { return x + y }(3, 4)
	fmt.Printf("a = %+v\n", a)

	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Printf("sum = %+v\n", sum)
	}()

	fn := func() {
		fmt.Println("hello")
	}
	fn()

	fmt.Println("匿名函数加法求和", func(x, y int) int { return x + y }(3, 4))

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

func list(nums ...int) {
	fmt.Println(nums)
}
