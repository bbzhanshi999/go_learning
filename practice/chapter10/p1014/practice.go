package main

import "fmt"

//为 int 定义一个别名类型 `Day`，定义一个字符串数组它包含一周七天的名字，为类型 `Day` 定义 `String()` 方法，它输出星期几的名字。使用 `iota` 定义一个枚举常量用于表示一周的中每天（MO、TU...）。

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SE
	SU
)

var Weeks []string = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

type Day int

func (d Day) String() string {
	return Weeks[d]
}

func main() {
	fmt.Println(MO)
	fmt.Println(SU)

	var th Day = 19
	fmt.Println(th)

	//Monday
	// Sunday
	// %!v(PANIC=String method: runtime error: index out of range)
}
