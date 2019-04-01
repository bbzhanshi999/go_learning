package main

import "fmt"

func main() {
	// p72()
	p73()
}

//练习7.1：array_value.go: 证明当数组赋值时，发生了数组内存拷贝。
func p71() {
	arr1 := [...]int{1, 2, 3, 4, 5}

	arr2 := arr1

	fmt.Printf("arr1 address is %p\n", &arr1)
	fmt.Printf("arr2 address is %p\n", &arr2)
}

//练习7.2：for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。
func p72() {
	var arr1 [16]int

	for i, v := range arr1 {
		v = i
		println(v)
	}
}

// 练习7.3：fibonacci_array.go: 通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。
func p73() {
	var fiboArr [50]int

	fibonacciArray(0, &fiboArr)

	for _, v := range fiboArr {
		println(v)
	}
}

func fibonacciArray(idx int, source *[50]int) {
	if idx == 0 {
		source[idx] = 1
	} else if idx == 1 {
		source[idx] = 1
	} else {
		source[idx] = source[idx-1] + source[idx-2]
	}
	idx += 1
	if idx < 50 {
		fibonacciArray(idx, source)
	}
}
