package main

import (
	"fmt"
	"os"
)

func main() {
	//不会打印该方法
	defer fmt.Println("!")

	os.Exit(3)
}
