package main

import (
    "fmt"
)

func main() {
   /* s := make([]string, 10)
    for i := 0; i < 3; i++ {
        s[i] = "tcy"
    }
    fmt.Println(strings.Join(s, "最棒")) //tcy最棒tcy最棒tcy最棒最棒最棒最棒最棒最棒最棒*/
    /*var b strings.Builder
    b.Grow(3)
    b.WriteString("1")
    for i := 0; i < 10; i++ {
        b.WriteString(strconv.Itoa(i))
    }
    fmt.Println(b.String())
    fmt.Println(b.Len())*/

    s := "1111"
    fmt.Println(&s)

    a := []byte(s)
    fmt.Printf("%p", a)
}