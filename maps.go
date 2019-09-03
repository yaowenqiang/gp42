package main

import (
	"fmt"
)

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
}
