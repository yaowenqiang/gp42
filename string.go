package main
import "fmt"
func main() {
  const s = "go语言"
  for i, r := range s {
    fmt.Printf("%#U    : %d\n", r, i)
  }
}
