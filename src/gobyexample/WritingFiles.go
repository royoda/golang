package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkw(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//写入文件方式一,将字节数组的值写入文件中，文件不存在则自动创建,每次写入都会覆盖
	d1 := []byte("hello\nworld\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	checkw(err)

	//创建文件
	//f, err := os.Create("/tmp/dat2")
	//checkw(err)
	//使用openfile加指定参数可以追加内容到指定文件中，os.O_APPEND写入数据时追加，os.O_RDWR以读写方式打开文件
	//os.Open()//return OpenFile(name, O_RDONLY, 0)从源码可以知道os.Open方法内部调用了os.OpenFile方法，且以只读的方式打开
	f, err := os.OpenFile("/tmp/dat", os.O_APPEND|os.O_RDWR, 0666)
	checkw(err)

	defer f.Close()
	b1 := make([]byte, 10)
	r1, err := f.Read(b1) //
	checkw(err)
	fmt.Println("r1:", string(b1[:r1]))

	//写入文件方式二，
	d2 := []byte{97, 98, 115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkw(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//写入文件方式三
	n3, err := f.WriteString("writes\n")
	checkw(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	//写入文件方式四
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkw(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
