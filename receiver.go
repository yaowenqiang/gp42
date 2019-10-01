// Package main provides ...
package main

import (
	"fmt"
)

type A struct {
	Face int
}

func (a A) f() {
	fmt.Println("hi, ", a.Face)
}
func main() {
	a := new(A)
	a.f()
}
