package main

import (
	"fmt"
	"os/exec"
	"unicode/utf8"
)

/**
    *@Description TODO
    *@Date 2021/4/29 4:42 下午
**/

type LogLoginTypeDao struct {
	CurrTableName string
	MapTableName  map[int]string
}

func main() {
	//fmt.Println(time.Now().Add(time.Duration(0) * time.Second).Format("200601"))
	str := "abcdef 世界"
	var n = 1
	for i, v := range str {
		if v < 0 {

		}
		fmt.Printf("%T   %T\n", str[i], v)
		n++
	}
	i, _ := utf8.DecodeRuneInString(str)

	fmt.Printf("%d %d %d\n", len(str), i, utf8.RuneCountInString(str))

	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("%p\n", arr)
	arr = append(arr[:2], arr[:len(arr)-1]...)
	fmt.Println(arr)
	fmt.Printf("%p\n", arr)
	arr = append(arr, 7)
	fmt.Printf("%p %d\n", arr, cap(arr))
	arr = append(arr, 7)
	fmt.Printf("%p %d\n", arr, cap(arr))
	fmt.Println(0i)
	fmt.Printf("%T\n", '\000')

	cmd := exec.Command("pwd")
	s, err := cmd.Output()
	fmt.Println(string(s), err)
}
