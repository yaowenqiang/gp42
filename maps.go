package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{"a": "1", "b": "2", "c": "3"}

	fmt.Println(map3)
	fmt.Println(map2)
	fmt.Println(map1)

	x := map[string]string{"tow": "2"}

	if _, ok := x["tow"]; !ok {
		fmt.Println("No entry")
	}
	fmt.Println(map3["a"])
	delete(map3, "a")
	fmt.Println(map3)

	switch a := 1; {
	case a == 1:
		fmt.Println("The integer was == 1")
		fallthrough
	case a == 2:
		fmt.Println("The integer was == 2")
	case a == 3:
		fmt.Println("The integer was == 3")
		fallthrough
	case a == 4:
		fmt.Println("The integer was == 4")
	case a == 5:
		fmt.Println("The integer was == 5")
		fallthrough
	default:
		fmt.Println("default case")
	}

	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received  ", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("send ", 12, "to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	case <-time.After(time.Second * 3): //超时退出
		fmt.Println("request time out")
	}
	for i, j := 0, 0; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %+v\n", i)
	}

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for j := 0; j < 5; j++ {
		for i := 0; i < 50; i++ {
			if i > 5 {
				break
			}
			fmt.Printf("i = %+v\n", i)
		}
	}

	data := []field{{"one"}, {"tow"}, {"three"}}

	for _, v := range data {
		go v.print()
	}

	time.Sleep(3 * time.Second)

	data2 := []string{"one", "tow", "three"}

	for _, v := range data2 {
		go func(in string) {
			fmt.Printf("in = %+v\n", in)
		}(v)
	}

	time.Sleep(3 * time.Second)
}
