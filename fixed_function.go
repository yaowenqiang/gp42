// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	s := []string{"James", "Jasmine"}
	Greeting(s...)
}

func Greeting(who ...string) {
	for k, v := range who {
		fmt.Println(k, v)
	}
}
