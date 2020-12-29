package main

import (
	"os"
)

//一个panic通常意味着出事意外错误。
func main() {
	//panic("a proablem")

	//创建文件,如果文件存在则覆盖
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
