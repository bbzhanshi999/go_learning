package main

import (
	"fmt"
	"strings"
)

//1）stack_arr.go：使用长度为 4 的 int 数组作为底层数据结构。
// 3）使用常量 LIMIT 代替上面表示元素个数的 4 重新实现上面的 1）和 2），使它们更具有一般性。

// 一个新栈中所有格子的值都是 0。

// 将一个新值放到栈的最顶部一个空（包括零）的格子中，这叫做`push`。

// 获取栈的最顶部一个非空（非零）的格子的值，这叫做`pop`。
// 现在可以理解为什么栈是一个后进先出（LIFO）的结构了吧。

// 为栈定义一个`Stack` 类型，并为它定义 `Push` 和 `Pop` 方法，再为它定义 `String()` 方法（用于调试）输出栈的内容，比如：`[0:i] [1:j] [2:k] [3:l]`。

const LIMIT = 3

type Stack [LIMIT]int

func (s *Stack) Push(t int) {

	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			s[i] = t
			return
		}
		if i == len(s)-1 {
			s[i] = t
			return
		}
	}
}

func (s *Stack) Pop() int {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 {
			var v int
			v, s[i] = s[i], 0
			return v
		}
	}
	return 0
}

func (s *Stack) String() string {

	strArr := make([]string, len(s))

	for i, v := range s {
		strArr[i] = fmt.Sprintf("[0:%d]", v)

	}
	return strings.Join(strArr, " ")
}

func main() {
	s := new(Stack)
	s.Push(10)
	//fmt.Println(s.Pop())

	s.Push(20)
	//fmt.Println(s.Pop())

	s.Push(30)
	//fmt.Println(s.Pop())

	s.Push(40)
	//fmt.Println(s.Pop())

	s.Push(50)
	//fmt.Println(s.Pop())

	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
}
