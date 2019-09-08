package main

import (
	"errors"
	"fmt"
)

func main() {
	f, err := fmt.Println(Sqrt(-0.1))
	if f == 1 {
		fmt.Printf("f = %+v\n", f)
	} else {
		panic(err)
	}

}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("Math - square root of negative number, the number is %d", f)
		return 0, fmt.Errorf(" the square root of negative number %g", f)
	}
	return 1, errors.New("Math - square root of negative number ")

}
