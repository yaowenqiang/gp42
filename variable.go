package main
import (
"fmt"
 "log"
 "time"
// _ "log"
// _ "time"
)

func main() {
  var (
    a string
  )
  a = "123"
  fmt.Printf("%s\n", a)

  d, e, f := 1,2,4
  _, g := 5,6
  fmt.Printf("g =  %d\n", g)

  fmt.Printf("%d %d %d", d, e, f)

  var _ = log.Println

  _ = time.Now

  x := [3]int{1,2,3}
  func(arr *[3]int) {
    (*arr)[0] = 7
    fmt.Println(arr)
  }(&x)
  fmt.Println(x)


  //ar x = nil

  //_ = x

  //var m map[string]int
  //m["one"] = 1
  var str string = ""
  fmt.Printf("%s", str)
  const PI = 3.14
  fmt.Printf("%f", PI)

  const (
    b = iota
    c = iota
    h = iota
  )
  fmt.Printf("\nb: %d\n", b)
  fmt.Printf("c: %d\n", c)
  fmt.Printf("h: %d\n", h)

  const bb = 5;
  var intVar int = bb
  var int32Var int32 = bb;
  var float64Var float64 = bb;
  var complex64Var complex64 = bb;
  fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)



}

