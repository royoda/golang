package main

import (
	"fmt"
)

func main() {
	/*currentTime := time.Now()
	currentTime1 := currentTime.Format("2006-01-02") + " 23:59:59"
	oldTime := currentTime.AddDate(0, 0, -1).Format("2006-01-02") + " 00:00:00"

	fmt.Println(currentTime1)
	fmt.Println(oldTime)*/

	// 分页获取数据逻辑
	pageSize := 100
	page := 1
	for i := 1000; i > 0; i = i - pageSize {
		fmt.Println(page)
		fmt.Println(pageSize)
		page = page + 1
	}
}
