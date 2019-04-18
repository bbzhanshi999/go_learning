package stack

import (
	"fmt"
	"strings"
)

// 一个新栈中所有格子的值都是 0。

// 将一个新值放到栈的最顶部一个空（包括零）的格子中，这叫做`push`。

// 获取栈的最顶部一个非空（非零）的格子的值，这叫做`pop`。
// 现在可以理解为什么栈是一个后进先出（LIFO）的结构了吧。
//2）stack_struct.go：使用包含一个索引和一个 int 数组的结构体作为底层数据结构，索引表示第一个空闲的位置。
//3）使用常量 LIMIT 代替上面表示元素个数的 4 重新实现上面的 1）和 2），使它们更具有一般性。

const LIMIT = 5

type Stack struct {
	index int
	arr   [LIMIT]int
}

func (s *Stack) Push(t int) {

	for i := 0; i < len(s.arr); i++ {
		if s.arr[i] == 0 || i == len(s.arr)-1 {
			s.arr[i] = t
			s.index = i + 1
			return
		}
	}
}

func (s *Stack) Pop() int {

	for i := len(s.arr) - 1; i >= 0; i-- {
		if s.arr[i] != 0 {
			var v int
			v, s.arr[i] = s.arr[i], 0
			if s.index != 0 {
				s.index -= 1
			}
			return v
		}
	}
	return 0
}

func (s *Stack) String() string {

	strArr := make([]string, len(s.arr))

	for i, v := range s.arr {
		strArr[i] = fmt.Sprintf("[0:%d]", v)

	}
	return strings.Join(strArr, " ")
}
