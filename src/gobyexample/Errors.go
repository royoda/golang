package main

import (
	"errors"
	"fmt"
)

/**
自定义程序错误报错提示和使用默认的情况
*/
//使用默认的errors
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can`t work with 42")
	}
	return arg + 3, nil
}

//使用自定义的errors
type arrError struct {
	arg  int
	prob string
}

func (e *arrError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//给指针地址赋值，并返回
		return -1, &arrError{arg, "can`t wor with 42"}
	}
	return arg + 3, nil
}
func main() {
	for _, v := range []int{7, 42} {
		if r, e := f1(v); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, v := range []int{7, 42} {
		if r, e := f2(v); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	_, e2 := f2(42)
	//f2方法执行返回拿到了e2的地址，ae,ok := e2.(*arrError) 返回ae结构体，ok表示是否指针指向的位置是否有值，
	//有值即正确
	//就返回地址对应的值
	if ae, ok := e2.(*arrError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
