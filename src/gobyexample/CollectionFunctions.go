package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	//i 是索引，v是值
	for i, v := range vs {
		//判断值是否等于传送过来的参数，相等就返回索引值
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
func ToNums(strs []string, f func(v string) int32) []int32 {
	//使用该方法创建切片，理解：*动态数组
	numarr := make([]int32, len(strs))
	for i, v := range strs {
		numarr[i] = f(v)
	}
	return numarr
}
func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear")) //2

	fmt.Println(Include(strs, "grape")) //false

	//验证数组中是否有p开头的数据，有一个就满足
	fmt.Println(Any(strs, func(v string) bool {
		//验证是否以p开头
		return strings.HasPrefix(v, "p")
	}))
	//验证数组中是否有p开头的数据，需要全部满足
	fmt.Println(All(strs, func(v string) bool {
		//验证是否以p开头
		return strings.HasPrefix(v, "p")
	}))

	//验证数组中数据包含e
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	//将数组的数据转为大写
	fmt.Println(Map(strs, strings.ToUpper))

	//将数组的数据转为数字
	fmt.Println(ToNums(strs, func(v string) int32 {
		var sum int32 = 0
		for _, c := range v {
			sum += c
		}
		return sum
	}))

}
