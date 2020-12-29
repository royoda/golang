package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkfile(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//创建文件
	/*file,err :=os.Create("/tmp/dat")
	checkfile(err)
	defer file.Close()

	//写文件方式一
	w3, err := file.WriteString("writes\n")//返回写入的字节数
	checkfile(err)
	fmt.Printf("wrote %d bytes\n", w3)*/

	/******************/
	//读文件方式一
	dat, err := ioutil.ReadFile("/tmp/dat")
	checkfile(err)
	fmt.Println("读取方式一")
	fmt.Print(string(dat))

	//打开文件
	f, err := os.Open("/tmp/dat")
	checkfile(err)
	//读文件方式二
	b1 := make([]byte, 5) //设置字节缓冲，用于存储读取到的文字
	n1, err := f.Read(b1)
	checkfile(err)
	fmt.Println("读文件方式二")
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//读文件方式三,寻找文件中第六个字节的位置开始,从指定位置开始读取
	o2, err := f.Seek(6, 0)
	checkfile(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2) //读取两个字节
	checkfile(err)
	fmt.Println("读文件方式三")
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	//读文件方式四
	o3, err := f.Seek(6, 0)
	checkfile(err)
	b3 := make([]byte, 2)
	//设置读取量ReadAtLeast
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkfile(err)
	fmt.Println("读文件方式四")
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//读文件方式五
	_, err = f.Seek(0, 0)
	checkfile(err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(10) //参数超出文件字节数会报错
	checkfile(err)
	fmt.Println("读文件方式五")
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
