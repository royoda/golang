package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//找到p 开头，ch结尾，中间由小写字母组成
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//定义正则表达式
	r, _ := regexp.Compile("p([a-z]+)ch")
	//匹配字符串
	fmt.Println(r.MatchString("peach"))
	//返回第一个匹配的结果，没有则返回空字符
	fmt.Println(r.FindString("punch peach punch"))
	//返回结果集的的开始索引和结束索引
	fmt.Println(r.FindStringIndex("peach punch"))
	//返回两个结果集一个是p([a-z]+)ch，另一个是([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//返回两个结果集的的开始索引和结束索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//返回所有匹配的结果集，数字是返回多少个结果
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	//返回匹配结果集中两个结果集一个是p([a-z]+)ch，另一个是([a-z]+)
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	//返回所有匹配的结果集，数字是返回多少个结果
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	//匹配字节数组
	fmt.Println(r.Match([]byte("peach")))
	//更安全的使用正则表达式
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	//替换全部匹配结果集
	fmt.Println(r.ReplaceAllString("a peach punch pinch", "<fruit>"))
	//创建字节数组
	in := []byte("a peach")
	//使用函数
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
