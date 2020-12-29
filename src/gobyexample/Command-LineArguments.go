package main

import (
	"fmt"
	"os"
)

//要使用命令行参数，需要使用go build 构建
//go build Command-LineArguments.go 生成可执行文件
//./Command-LineArguments a b c d  模拟输入参数索引分别是0，1，2，3，4
//
func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
