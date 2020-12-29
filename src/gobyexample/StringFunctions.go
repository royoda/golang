package main

import (
	"fmt"
)

var p = fmt.Println

func main() {

	/*p("Contains:  ", str.Contains("test", "es"))
	p("Count:     ", str.Count("test", "t"))
	p("HasPrefix: ", str.HasPrefix("test", "te"))
	p("HasSuffix: ", str.HasSuffix("test", "st"))
	p("Index:     ", str.Index("test", "e"))
	p("Join:      ", str.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", str.Repeat("a", 5))
	p("Replace:   ", str.Replace("foo", "o", "0", -1))
	p("Replace:   ", str.Replace("foo", "o", "0", 1))
	p("Split:     ", str.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", str.ToLower("TEST"))
	p("ToUpper:   ", str.ToUpper("test"))
	p()

	p("Len: ", len("hello"))*/
	//字符串索引位置的字符asicc码,类型int8
	//var a = 8
	//fmt.Printf("%T\n",string(a))//%T输出类型
	p("Char:", "hello"[1])
}
