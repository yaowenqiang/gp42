package main

import (
    "fmt"
)

func main() {
    var x = []int{1,3,5,7,11}
    fmt.Printf("%v\n", x)
    var y = make([]int, 10,50)
    fmt.Printf("%v\n",y)
    a := [5]int{1,2,3,4,5}
    fmt.Printf("%v\n",a)
    t := a[1:3:5]
    fmt.Printf("%v\n",t)
    sli := make([]int, 5,10)
    fmt.Printf("切片 sli 长度和容量: %d, %d, \n", len(sli), cap(sli))
}

