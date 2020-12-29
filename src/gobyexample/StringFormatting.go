package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	//原格式输出
	fmt.Printf("%v\n", p)
	//如果是结构，则包含结构输出
	fmt.Printf("%+v\n", p)
	//打印改值源代码片段
	fmt.Printf("%#v\n", p)
	//打印值的类型
	fmt.Printf("%T\n", p)
	//
	fmt.Printf("%t\n", true)
	//十进制输出
	fmt.Printf("%d\n", 123)
	//二进制输出
	fmt.Printf("%b\n", 14)
	//打印整数对应的字符
	fmt.Printf("%c\n", 33)
	//十六进制输出
	fmt.Printf("%x\n", 456)
	//浮点数输出
	fmt.Printf("%f\n", 78.9)
	//科学计数法输出
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	//基本字符串打印
	fmt.Printf("%s\n", "\"string\"")
	//字符串原样输出
	fmt.Printf("%q\n", "\"string\"")
	//一个字符用两个数字输出,%x表示以十六进制形式输出整数。
	fmt.Printf("%x\n", "hex this")
	//打印指针地址
	fmt.Printf("%p\n", &p)
	//指定整数宽度6，默认情况向右对齐，不足的空间用空格填充
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	//和上面类似，不过这里用的是浮点数
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	//和上面类似，-号将会使齐向左对其
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//上述是控制数字，这里是控制字符串对其，默认向右对齐
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	//向左对齐
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	//格式化并返回字符串
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	// 写入文件中
	/*txt,_ := os.Create("/tmp/defer.txt")
	fmt.Fprintf(txt, "an %s\n", "error")*/

	//txt, err := os.OpenFile("/tmp/defer.txt", os.O_APPEND, 0666)
	/*fmt.Println(err)
	if err != nil {
		panic(err)
	}*/
	/*err := os.Chmod("/tmp/defer.txt", 0777)
	if err != nil {
		log.Println(err)
	}*/
	/*
		方式1：
		d1 := []byte("hello\ngo\n")
		err := ioutil.WriteFile("/tmp/defer.txt", d1, 0644)
		check(err)
	*/
	//方式2：
	f, err := os.OpenFile("/tmp/defer.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer f.Close()
	n3, err := f.WriteString("writes\n")
	//判断错误信息是否为空，为空则说明写入读取成功
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
