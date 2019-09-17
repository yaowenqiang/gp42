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

	const (
		c1 = imag(2i)
		c2 = len([10]float64{2})
		c3 = len([10]float64{c1})
		c4 = len([10]float64{imag(2i)})
		//c5 = len([1]float64{imag(z)})
	)

	fmt.Printf("c1 = %+v\n", c1)
	fmt.Printf("c2 = %+v\n", c2)
	fmt.Printf("c3 = %+v\n", c3)
	fmt.Printf("c4 = %+v\n", c4)
	var z complex128
	fmt.Printf("z = %+v\n", z)

}
