// Package main provides ...
package main

import "fmt"

type Point3D struct {
	x, y, z float64
}

func main() {
	//var pointer *Point3D

	var i *[4]int

	fmt.Printf("i = %+v\n", i)

	var a, b int = 20, 30

	var ptra *int
	var ptrb *int = &b

	ptra = &a

	fmt.Printf("a = %x\n", &a)
	fmt.Printf("b = %x\n", &b)

	fmt.Printf("ptra = %x\n", ptra)
	fmt.Printf("ptrb = %x\n", ptrb)

	fmt.Printf("*ptra = %d\n", *ptra)
	fmt.Printf("*ptrb = %d\n", *ptrb)

	//go run -gcflags -m pointer.go
	// 显示变量分配位置(堆或者栈)
}
