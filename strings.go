package main
import (
    "fmt"
    "unicode/utf8"
)
func main() {
    str := "Hello world! \nHello Gopher! \n"
    str2 := `Hello world! \nHello Gopher! \n`
    fmt.Println(str)
    fmt.Println(str2)

    str3 := `这个是中文`
    fmt.Println(str3)
    fmt.Println(len(str3))
    fmt.Println(utf8.RuneCountInString(str3))


}

