package main

import "fmt"

func main() {

}

func test1() {
FUCK:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j == 5 {
				break FUCK
			}
			println("in nested loop")
		}
		println("in root loop")
	}
}

func test2() {
	a := 1
	goto TARGET
	//b := 9 // 在正序标签之间声明变量会编译错误，很好理解，因为你都跳到了标签，那么标签之间的代码是不会执行的，无论在标签之后的代码是否使用该声明变量
TARGET:
	b := 9
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
