// Package main provides ...
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	Age  int
}

func main() {
	var a int = 10
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	fmt.Printf("v = %+v\n", v)
	fmt.Printf("t = %+v\n", t)

	fmt.Printf("t.Kind() = %+v\n", t.Kind())

	fmt.Printf("v.Type() = %+v\n", v.Type())

	fmt.Printf("reflect.ValueOf(&a).Elem() = %+v\n", reflect.ValueOf(&a).Elem())

	seta := reflect.ValueOf(&a).Elem()

	seta.SetInt(1000)

	fmt.Printf("seta = %+v\n", seta)
	fmt.Printf("a = %+v\n", a)

	a = 3
	fmt.Printf("a = %+v\n", a)
	fmt.Printf("seta = %+v\n", seta)

	var b [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("reflect.TypeOf(b).Elem() = %+v\n", reflect.TypeOf(b).Elem())

	var Pupil Student = Student{"joke", 18}

	p := reflect.ValueOf(Pupil)

	fmt.Printf("p.Type() = %+v\n", p.Type())
	fmt.Printf("p.Kind() = %+v\n", p.Kind())
	fmt.Printf("Pupil.name = %+v\n", Pupil.name)

	setStudent := reflect.ValueOf(&Pupil).Elem()

	//setStudent.Field(0).SetString("Mike") ,字段没有导出，不能修改
	setStudent.Field(1).SetInt(20)

	fmt.Printf("setStudent = %+v\n", setStudent)
	fmt.Printf("Pupil = %+v\n", Pupil)

}
