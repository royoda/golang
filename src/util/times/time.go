package main

import "fmt"
import "github.com/golang-module/carbon"
func main()  {
	cur := carbon.Now()
	yesterdayStartTime := cur.Yesterday().StartOfDay()
	endTime := yesterdayStartTime
	for endTime != cur.Yesterday().EndOfDay() {

		endTime = yesterdayStartTime.AddMinutes(9).EndOfMinute()

		fmt.Println(yesterdayStartTime.ToDateTimeString(),endTime.ToDateTimeString())


		yesterdayStartTime = yesterdayStartTime.AddMinutes(10)
	}

	//yesterdayStartTime := cur.AddMinutes(i * 10)
	//yesterdayEndTime := yesterdayStartTime.AddMinutes(10).ToDateTimeString()

	// 本分钟开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
	// 本分钟结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59
}