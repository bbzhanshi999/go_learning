package main

import (
	"fmt"
)

func main() {
	mf := map[int]func(int) int{
		1: func(factor int) int { return 1 * factor },
		2: func(factor int) int { return 2 * factor },
		3: func(factor int) int { return 3 * factor }}

	fmt.Println(mf)
	fmt.Println(mf[1](10)) //执行函数
	fmt.Println(mf[2](10))
	fmt.Println(mf[3](10))

}
