package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() float64
	perim() float64
}
type recit struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r recit) area() float64 {
	return r.height * r.width
}

func (r recit) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return c.radius * math.Pi * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) perim1() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := recit{width: 10, height: 10}
	c := circle{radius: 3}
	fmt.Println(c.perim1())
	measure(r)
	measure(c)
}
