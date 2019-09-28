// Package main provides ...
package main

import (
	"fmt"
)

type Namer interface {
	//Method1(param_list) return_type
	//Method2(param_list) return_type
	//	...
}
type A struct {
	Books int
}

type B interface {
	f()
}

/*
type Bad interface {
	Bad
}

type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

*/

func (a A) f() {
	fmt.Println("A:f()", a.Books)
}

type I int
type S interface {
	f()
}

func (i I) f() {
	fmt.Println("I.f() ", i)
}

/*
func (s S) f() {
	fmt.Println("S.f() ", s)
}
*/

func main() {
	//var i interface{} = 99
	//i = 44.09
	///fmt.Printf("i = %+v\n", i)
	//i = "ALL"
	//fmt.Printf("i = %+v\n", i)

	var a A = A{Books: 9}
	a.f()
	var b B = A{Books: 99}

	b.f()

	var i I = 199
	i.f()

	var b2 B = I(299)

	b2.f()

	/*
		var varS S
		varS = S("Tstring")

		if v, ok := varS.(T); ok {
			fmt.Println("varI 类型断言结果为: ", v)
			varS.f()
		}
	*/
	funcName("abc")
	funcName(123)

}

func funcName(a interface{}) string {

	//fmt.Println("Where are you ,Jonny?", a.(string))

	value, ok := a.(string)

	if !ok {
		fmt.Println("It's not ok for type string")
		return ""

	}

	fmt.Println("The value is : ", value)
	return value

}
