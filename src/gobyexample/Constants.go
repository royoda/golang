package main

import (
	"fmt"  //打印时需要的包
	"math" //函数类型的包
)

//常量,一旦赋值之后就不能在进行修改
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
