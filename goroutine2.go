// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 1; i < 10; i++ {
		time.Sleep(100 + time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}
