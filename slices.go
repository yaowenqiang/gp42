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
	data2 := get2()
	fmt.Println(len(data2), cap(data2), &data2[0])
	append_demo()

}

func get() []byte {
	raw := make([]byte, 10000)
	return raw[:3]
}

func get2() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func append_demo() {
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	fmt.Printf("s1 = %+v\n", s1)

	s2 := append(s1, 3, 5, 7)
	fmt.Printf("s3 = %+v\n", s2)

	s3 := append(s2, s0...)
	fmt.Printf("s3 = %+v\n", s3)

	s4 := append(s3[3:6], s3[2:]...)
	fmt.Printf("s4 = %+v\n", s4)
}
