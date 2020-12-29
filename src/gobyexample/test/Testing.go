package main

/**
golang 测试，测试需要两个文件，一个存放需要测试方法的xxx，一个存放test方法的编写xxx_test
*/
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetMaxValue(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
