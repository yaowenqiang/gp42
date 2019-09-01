package main

import (
	"fmt"
)

func main() {
	var x = []int{1, 3, 5, 7, 11}
	fmt.Printf("%v\n", x)
	var y = make([]int, 10, 50)
	fmt.Printf("%v\n", y)
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	t := a[1:3:5]
	fmt.Printf("%v\n", t)
	sli := make([]int, 5, 10)

	fmt.Printf("切片 sli 长度和容量: %d, %d, \n", len(sli), cap(sli))
	data := get()

	fmt.Println(len(data), cap(data), &data[0])

}

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}
