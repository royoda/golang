package main

import "fmt"

//访问迭代各种数组结构中的元素，遍历数组，切片，map
func main() {
	nums := []int{2, 3, 5}
	fmt.Println(nums)
	sum := 0
	//遍历数组，_ 代表空白符，否则num输出的是下标索引
	for _, num := range nums {
		fmt.Print(num, "\t")
		sum = +num
	}
	fmt.Println("\nsum", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Printf("%d\n", i)
		}
	}
	ksv := map[string]string{"a": "a1", "b": "b1"}
	//ksv["c"] = "c1"
	//遍历key，value
	for k, v := range ksv {
		fmt.Printf("%s ---> %s\n", k, v)
	}

	//变量map中的key
	for k := range ksv {
		fmt.Println("key:", k)
	}
	//遍历value
	for _, v := range ksv {
		fmt.Printf("value:%s\n", v)
	}
	//获取字符串中的字符下标索引和ascii码值
	for i, c := range "go" {
		fmt.Println("index:", i, "\tvalue:", c)
	}

}
