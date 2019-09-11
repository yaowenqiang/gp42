package main

import (
	"fmt"
	"log"
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

func protect(g func()) {
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g()
}

func main() {
	protect(test)
	div(10, 0)
	div(10, -1)
}
func test() {
	fmt.Printf(" hello world\n")
}
