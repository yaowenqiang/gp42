// Package main provides ...
package main

import (
	"fmt"
)

type Writer interface {
	/* TODO: add methods */
}

type Reader interface {
	/* TODO: add methods */
}

type Author struct {
	name string
	Writer
}

type Other struct {
	i int
}

func (a Author) Write() {
	fmt.Println(a.name, " Write.")
}

func (o Other) Write() {
	fmt.Println("Other  Write.")
}

func main() {
	Ao := Author{"Other", Other{99}}
	Ao.Write()

	Au := Author{name: "Hawking"}
	Au.Write()
}
