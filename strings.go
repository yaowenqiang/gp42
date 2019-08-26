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

    s := "去年今日此门中，人面桃花相映红"
    for k, v := range s {
        fmt.Printf("k: %d, v: %c == %d\n", k,v ,v)
    }

    s4 := "Beginning of the string" +
    "second part of the string"

    s5 := "Hel" + "lo, "
    s5  += "World!"
    fmt.Println(s4)
    fmt.Println(s5)

}

