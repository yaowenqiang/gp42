package main

import (
	"fmt"
)

func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常: %s\n", r)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}
}

func main() {
	div(10, 0)
	div(10, -1)
}
