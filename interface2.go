// Package main provides ...
package main

import (
	"fmt"
)

type I interface {
	f()
}

type T string

func (t T) f() {
	fmt.Println("T Method")
}

type Stringer interface {
	String() string
}

func main() {
	var varI I

	varI = T("Tstring")

	if v, ok := varI.(T); ok {
		fmt.Println("varI 类型断言结果为: ", v)
		varI.f()
	}

	var value interface{}

	switch str := value.(type) {
	case string:
		fmt.Println("value类型断言结果为 string", str)
	case Stringer:
		fmt.Println("value类型断言结果为 Stringer", str)
	default:
		fmt.Println("value类型断言不在上术类型之中")
	}

	value = "类型断言检查"

	str, ok := value.(string)
	if ok {
		fmt.Printf("value类型断言结果为 %T\n", str)
	} else {
		fmt.Println("value 不是string类型 %T\n")
	}

}
