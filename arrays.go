package main
import (
    "fmt"
    "github.com/kr/pretty"
)
func main() {
    var arrAge = [5]int{10,20,30,40,50}
    fmt.Printf("%v\n", arrAge)
    var arrName =  [5]string{1: "Jack", 3:"John"}
    fmt.Printf("%v\n", arrName)
    var arrcount = [4]int{500, 2:100}
    fmt.Printf("%v\n", arrcount)
    var arrLazy = [...]int{5,6,7,8,22}
    fmt.Printf("%v\n", arrLazy)
    var arrPack = [...]int{10,5:300}
    fmt.Printf("%v\n", arrPack)
    var arrRoom [20]int
    fmt.Printf("%v\n", arrRoom)
    var arrBed = new([20]int)
    fmt.Printf("%v\n",  arrBed)
    fmt.Printf("-----------\n")

    var arr1 = new([5]int)

    arr := arr1
    arr1[2] = 100
    fmt.Println(arr1[2], arr[2])

    var arr2 [5]int
    newarr := arr2
    arr2[2] = 100
    fmt.Println(arr2[2], newarr[2])
}


