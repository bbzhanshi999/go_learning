package main

import (
	"fmt"
)

func main() {
	typeSwitch2()
	typeSwtich3()
}

// swtich直接用表达式
func typeSwitch2() {
	target := 10

	switch {
	case target < 100:
		fmt.Printf("target %v <100\n", target)
		fallthrough //如果希望执行后续代码，需要加入fallthrough关键字
	case target < 20:
		fmt.Printf("target %v <20\n", target) //在go中，地一个case处理过后，第二个case即使符合条件也不处理了，除非使用fallthrough
	case target < 10:
		fmt.Printf("target %v <10\n", target)
	default:d
		fmt.Println("that is funny")
	}

}

func typeSwtich3() {
	target := 20

	switch result := calculate(target, 20); {
	case result < 0:
		println("less than")
	case result > 0:
		println("greater than")
	default:
		println("equal")
	}

	println(target)
}

func calculate(value int, compare int) int {
	if value < compare {
		value = 30 //并没有改变外面的target，因为并不是一个指针
		return -1
	} else if value > compare {
		return 1
	}
	return 0
}
