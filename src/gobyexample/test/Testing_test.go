package main

import (
	"fmt"
	"testing"
)

//testing.B 压力测试
func BenchmarkIntMin(b *testing.B) {

}
func TestIntMinBasic(t *testing.T) {

	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

//testing.T 功能测试
/*func TestIntMinTableDriven(t *testing.T) {

	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)//参数拼接
		//子测试
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}*/

func TestGetMaxValue(t *testing.T) {

	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{1, 2, 20, 4, 5, 6, 18}, 20},
		{[]int{1, 2, 3, 4, 5, 9}, 9},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.arr) //参数拼接
		t.Run(testname, func(t *testing.T) {
			ans := GetMaxValue(tt.arr)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
