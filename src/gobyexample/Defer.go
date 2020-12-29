package main

import (
	"fmt"
	"os"
)

func main() {

	//创建文件
	f := createFile("/tmp/defer.txt")
	//打开文件
	//f,_ := os.Open("/tmp/defer.txt")
	defer closeFile(f) //defer关键字会延迟执行该方法，即最后执行。该方法内
	//写入文件
	writeFile(f)
}

//创建文件
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

//写入文件
func writeFile(f *os.File) {
	//终端打印数据
	fmt.Println("writing")
	//写数据到文件中
	fmt.Fprintln(f, "dataasdasd")

}

//关闭文件
func closeFile(f *os.File) {

	//关闭方法
	err := f.Close()
	if err != nil {
		//%v 按默认格式输出
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		//程序退出
		os.Exit(1)
	}
}
