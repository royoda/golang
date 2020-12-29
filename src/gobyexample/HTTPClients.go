package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	//1创建文件
	f, _ := os.Create("/tmp/go.html")
	defer f.Close()
	for i := 0; scanner.Scan(); i++ {
		//fmt.Println(scanner.Text())
		//2写入文件
		f.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
