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
}
