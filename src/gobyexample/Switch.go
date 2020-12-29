package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")
	//第一种用法
	switch i {
	case 1:
		fmt.Println("i am one")
	case 2:
		fmt.Println("i am two")
	}
	fmt.Println(time.Now().Weekday())
	//第二种用法
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it`s a weekend")
	default:
		fmt.Println("it`s not a weekend but it`s weekday")
	}

	//第三种用法
	whatAmI := func(i interface{}) {
		//这种switch和val.(type)配合的语法是特有的，在switch以外的任何地方都不能使用类似于val.(type)这种形式。
		switch temp := i.(type) {
		case bool:
			fmt.Println("i am is bool")
		//case int:
		//	fmt.Println("i am is int") // 注释int走默认方法
		case string:
			fmt.Println("i am is string")
			//println("这也能输出")
		default:
			fmt.Println("dont`t know type", temp)
		}
	}
	whatAmI("hello")
	whatAmI(12)
	whatAmI(true)

}
