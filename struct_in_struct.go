// Package main provides ...
package main

import (
	"fmt"
)

type Human struct {
	name   string
	Gender string
	Age    int
	string
}

type Student struct {
	Human
	Room int
	int
}

func main() {
	stu := new(Student)

	stu.Room = 102
	stu.Human.name = "Titan"
	stu.Gender = "男"
	stu.Human.Age = 14
	stu.Human.string = "Student"

	fmt.Printf("stu = %+v\n", stu)

	stud := Student{Room: 102, Human: Human{"Hawking", "男", 14, "Monitor"}}
	fmt.Printf("stud = %+v\n", stud)
}
