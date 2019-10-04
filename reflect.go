// Package main provides ...
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

func main() {
	var a int = 50

	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	fmt.Printf("v = %+v\n", v)
	fmt.Printf("t = %+v\n", t)

	var b [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("TypeOf(b) = %+v\n", reflect.TypeOf(b))
	fmt.Printf("TypeOf(b).Kind() = %+v\n", reflect.TypeOf(b).Kind())
	fmt.Printf("TypeOf(b).Elem() = %+v\n", reflect.TypeOf(b).Elem())

	var Pupil Student

	p := reflect.ValueOf(Pupil)

	fmt.Printf("p.Type() = %+v\n", p.Type())
	fmt.Printf("p.Kind() = %+v\n", p.Kind())

}
