// Package main provides ...
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	go singalListen()
	for {
		time.Sleep(10 * time.Second)
	}

}

func singalListen() {
	c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGUSR2)
	signal.Notify(c, os.Interrupt)

	for {
		s := <-c
		fmt.Printf("get singal s = %+v\n", s)
	}
}
