package main

import (
	"fmt"
)

func main() {

}

func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return 1
		}
	}

	return
}
