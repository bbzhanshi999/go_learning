package main

import (
	"fmt"

	util "./pack"
)

func main() {
	util.Hello()
	fmt.Println("dsdfdsf4")

	f1(f2(1, 2))                            //在函数多返回值的情况下，假定外层某一函数的入参类型和个数顺序与另一个函数的返回值正好一致的话，就可以这样来编写代码
	var bf binOp = func(a int, b int) int { //实现一个函数
		return a + b
	}
	bf(1, 2)

}

func f1(a, b, c int) {
	fmt.Printf("a:%v b:%v c:%v ", a, b, c)
}

func f2(a, b int) (int, int, int) {
	return a, b, a + b
}

type binOp func(int, int) int //声明一个函数类型的别名
