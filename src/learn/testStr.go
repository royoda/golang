package main

import (
	"fmt"
)

/**
    *@Description TODO
    *@Date 2021/3/23 3:23 下午
**/
func main() {
	/*fieldKey := fmt.Sprintf("%s", "NaN")
	fieldKeyFloat, err := strconv.ParseFloat(fieldKey, 64)
	if fieldKeyFloat < 0 {
		fmt.Println(fieldKeyFloat)
	}
	fmt.Println(fieldKeyFloat)
	fmt.Println(err)*/
	/*nums := []int{3,2,2,3}
	end := removeElement(nums, 2)
	nums = nums[:end]
	fmt.Println(nums)*/

	s := strStr("aaaa", "bba")
	fmt.Println(s)
}

func removeElement(nums []int, val int) int {
	st := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			tmp := nums[st]
			nums[st] = nums[i]
			nums[i] = tmp
			st -= 1
		}
	}
	return st + 1
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	st := 0
	index := 0
	ed := 0
	l := len(haystack)
	for len(haystack) > index {
		for i := index; i < l; i++ {
			if haystack[i] != needle[st] {
				index += 1
				st = 0
				ed += 1
				break
			} else {
				st = st + 1
				if st == len(needle) {
					return ed
				}
			}
		}
	}

	return -1
}
