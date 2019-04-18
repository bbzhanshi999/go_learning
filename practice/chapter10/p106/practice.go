package main

import "fmt"

//定义结构体 `employee`，它有一个 `salary` 字段，给这个结构体定义一个方法 `giveRaise` 来按照指定的百分比增加薪水。

type employee struct {
	salary float64
}

func (ep *employee) giveRaise(percent float64) {
	ep.salary = ep.salary * (1 + percent)
}

func main() {
	ep := employee{10000.00}
	ep.giveRaise(0.3)
	fmt.Println(ep)
}
