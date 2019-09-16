package main

import (
	"fmt"
)

type S struct {
	a int
	b float64
}

func main() {
	s := make([]int, 10, 100)
	fmt.Printf("s = %+v\n", s)
	//s = [0 0 0 0 0 0 0 0 0 0]
	s1 := make([]int, 1e3)
	fmt.Printf("s1 = %+v\n", s1)

	//s2 := make([]int, 10, 0)
	//fmt.Printf("s2 = %+v\n", s2)
	//illegal: len(s) > cap(s)

	//s3 := make([]int, 1<<63)
	//fmt.Printf("s3 = %+v\n", s3)
	//len argument too large in make([]int)

	c := make(chan int, 10)
	fmt.Printf("c = %+v\n", c)

	m := make(map[string]int, 100)
	fmt.Printf("m = %+v\n", m)
	sss := new(S)
	fmt.Printf("new(S) = %+v\n", new(S))
	fmt.Printf("sss = %+v\n", sss)

}
