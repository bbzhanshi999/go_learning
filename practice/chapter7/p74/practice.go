package main

import "fmt"

//练习 7.4： fobinacci_funcarray.go: 为练习 7.3 写一个新的版本，主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片。
func main() {
	fmt.Println(fibonacciSlice(50))
}

func fibonacciSlice(length int) (slice []int) {
	slice = make([]int, length)
	slice[0], slice[1] = 1, 1
	for i := 2; i < len(slice); i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return
}
