package main

import (
	"fmt"
	"time"
)

func main() {
	var date time.Time = time.Now()

	println(date.Day())
	println(date.Minute())
	println(date.Second())
	println(date.Format("2006-01-02 15:04:05")) //按照格式化输出日期字符串,必须按照日期 2006年一月二日十五点零四分零五秒来表示

	fmt.Printf("%v-%02d-%02v\n", date.Year(), date.Month(), date.Day())

	//时间计算
	var week int64 = 60 * 60 * 24 * 7 * 1e9
	fmt.Printf("after a week time is: %v\n", date.Add(time.Duration(week)))   //加一周
	fmt.Printf("before a week time is: %v\n", date.Add(time.Duration(-week))) //减一周

	someday, _ := time.Parse("2006-01-02 15:04:05", "2018-10-10 00:00:00") //将字符串按照指定格式转换成time.Time

	println(someday.Format("2006-01-02 15:04:05"))
	println(date.After(someday))

}
