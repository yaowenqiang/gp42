// Package main provides ...
package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 1024)
	f, err := os.Open("./tt.txt")
	_, err = f.Read(b)

	f.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("string(b) = %+v\n", string(b))
}
