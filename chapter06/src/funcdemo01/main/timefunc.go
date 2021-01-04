package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now = %v now type = %T\n",now, now)

	fmt.Printf("Year=%v\n", now.Year())
	fmt.Printf("Month=%v\n", now.Month()) //int(now.Month())月份可以用int强转为数字
	fmt.Printf("Day=%v\n", now.Day())
	fmt.Printf("Hour=%v\n", now.Hour())
	fmt.Printf("Minute=%v\n", now.Minute())
	fmt.Printf("Second=%v\n", now.Second())

	fmt.Printf("Current 年月日 %d-%d-%d :%d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("Current 年月日 %d-%d-%d :%d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr = %v", dateStr)

	//格式化日期时间的第二种方式
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	////每隔一秒中打印一个数字，打印100时就退出
	//stopTime := 0
	//for {
	//	stopTime++
	//	fmt.Println(stopTime)
	//	//sleep
	//	time.Sleep(time.Millisecond * 100)
	//	//time.Sleep(time.Second)
	//	if stopTime >= 10 {
	//		break
	//	}
	//}
	//Unix&UnixNano的使用
	fmt.Printf("unix时间戳=%v, unixnano时间戳=%v\n", now.Unix(), now.UnixNano())




}
