package main

import "fmt"

func main()  {
    m := make(map[int]string, 0)
    m[1] = "1"
    m[2] = "2"
    m[3] = "3"
    m[4] = "4"
    for i, s := range m {
        fmt.Println(i,s)
    }
    delete(m,1)
    for i, s := range m {
        fmt.Println(i,s)
    }
}
