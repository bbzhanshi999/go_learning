package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getx2x3(num)
	fmt.Printf("x2 %v x3 %v \n", numx2, numx3)
	numx2, numx3 = getx2x3_1(num)
	fmt.Printf("x2 %v x3 %v \n", numx2, numx3)
}

func getx2x3(num int) (int, int) {
	return num * 2, num * 3
}

//命名返回参数在代码执行期就已经有了初始零值，即使不赋值，也能返回零值回去，而return语句也不需要列出返回的变量
func getx2x3_1(num int) (x2 int, x3 int) {
	x2 = num * 2
	x3 = num * 3
	return
}
