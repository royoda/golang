package main

import "fmt"

//map 关联数据类型
//make(map[key-type]val-type)
func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("maps：", m)

	//返回的是map键k1的值和true(说明存在)，false说明不存在
	v2, ok := m["k1"]
	fmt.Println("是否存在key：", ok, "对应的value：", v2)
	//当key=k1存在时打印value
	if v2, ok := m["k1"]; ok {
		fmt.Println("value:", v2)
	}
	//返回的是map中k1的值
	v1 := m["k1"]
	fmt.Println(v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println(m)

	//判断key是否存在
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
