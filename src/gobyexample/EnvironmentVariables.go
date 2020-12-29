package main

import (
	"fmt"
	"os"
	"strings"
)

//获取计算机的环境变量值
func main() {
	//os.Setenv("FOO", "1")
	//fmt.Println("FOO:", os.Getenv("FOO"))
	//需要大写
	fmt.Println("BAR:", os.Getenv("PATH"))

	fmt.Println()
	for _, e := range os.Environ() {
		fmt.Println(e)
		//截取字符串  例子：HOME=/Users/royod
		pair := strings.SplitN(e, "=", 2)
		//pairsub := strings.SplitN(pair[1], ":",-1)//获取环境变量值中以:分割获取，-1表示获取全部
		fmt.Println(pair[0])
		//fmt.Println(len(pairsub))
	}
}
