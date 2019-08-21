package main
import "fmt"
func main() {
  const s = "go语言"
  for i, r := range s {
    fmt.Printf("%#U    : %d\n", r, i)
  }

  var ch int = '\u0041'
  var ch1 int = '\u0382'
  var ch2 int = '\u0010'
  fmt.Printf"%#u", ch)
}
