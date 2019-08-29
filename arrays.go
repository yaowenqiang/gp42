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

    var arr3 = [...][5]int{{10,20},{30,40}}

    fmt.Printf("%# v", pretty.Formatter(arr3))

    type myType struct {
        a,b int
    }
    var x = []myType{{1,2}, {3,4},{5,6}}

    fmt.Printf("%# v\n", pretty.Formatter(x))
    fmt.Printf("%v\n", x)


    b := [...][5]int{{10,20},{30,40}, {30,40,50,60}}
    fmt.Println(b[1][1], len(b))
    fmt.Printf("%# v\n", pretty.Formatter(b))


    var arrAge2 = [5]int{1,2,3,4,5}
    for i, v := range arrAge2 {
        fmt.Printf("%d 的年龄: %d\n", i, v)
    }


}


