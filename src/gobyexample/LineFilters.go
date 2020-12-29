package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//scanner 输入
func main() {
	//os.Stdin 标准输出
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//用户输入的结果，转成大写
		ucl := strings.ToUpper(scanner.Text())
		//输出结果
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
