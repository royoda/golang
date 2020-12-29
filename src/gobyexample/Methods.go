package main

import "fmt"

type rect struct {
	width, height int
}

/**
方法定义方式，使用该方式时需要给对象赋值才能在通过指针去访问方法
*/
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

}
