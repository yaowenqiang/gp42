// Package main provides ...
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	Age  int `json:"years"`
}

func main() {
	var Pupil Student = Student{"Joke", 20}

	setStudent := reflect.ValueOf(&Pupil).Elem()

	sSAge, _ := setStudent.Type().FieldByName("Age")
	fmt.Printf("sSAge = %+v\n", sSAge)
	fmt.Printf("sSAge.Tag.Get('json') = %+v\n", sSAge.Tag.Get("json"))

}
