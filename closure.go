// Package main provides ...
package main

import (
	"fmt"
)

var G int = 7

func main() {
	y := func() int {
		fmt.Printf("G: %d, G的地址: %p\n", G, &G)
		G += 1
		return G
	}

	fmt.Printf("y() = %+v\n", y())
	fmt.Printf("y() = %+v\n", y())
	fmt.Printf("y() = %+v\n", y())

	z := func() int {
		G += 1
		return G
	}()

	fmt.Println(z, &z)
	fmt.Println(z, &z)
	fmt.Println(z, &z)

	var f = N()
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)

	var f1 = N()

	fmt.Println(f1(2), &f1)

}

func N() func(int) int {
	var i int
	return func(d int) int {
		fmt.Printf("i:%d, i的地址: %p\n", i, &i)
		i += d
		return i
	}
}
