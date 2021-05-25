package main

import "fmt"

//切片，切片和数组的区别：1，不提前声明数组的容量的方式即是切片。2，切片可以使用append方式追加数据
func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println(len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("app", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpoy", c)

	l := s[2:5] //截取索引2-5的数据[2，5）
	fmt.Println("l1", l)
	//说明值是指向地址，会因为原数组的值变动而变动
	s[2] = "2"
	fmt.Println("l1", l)

	l = s[:5]
	fmt.Println("s", s)
	fmt.Println("l2", l)
	l = s[2:]
	fmt.Println("l3", l)
	t := []string{"g", "h", "f"}
	fmt.Println("t", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//声明切片
	var arrslice = make([]int, 3)
	arrslice = append(arrslice, 12)
	fmt.Println("arrslice:", arrslice)

	//删除元素
	var arrRe = []int{1, 2, 3, 4, 5, 6, 7}
	res := remove(arrRe, 3)
	fmt.Println(res)

}

// 删除元素
func remove(slice []int, i int) []int {
	//copy(slice[i:], slice[i+1:])
	copy(slice[:i], slice[i+1:])

	return append(slice[:i], slice[i+1:]...)
}
