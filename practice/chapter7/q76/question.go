package main

import "fmt"

//演示可变参数的用法以及append的用法
func main() {
	var arr []int = []int{1, 2, 3}
	arr = appendElement(arr, 5, 6, 7)
	fmt.Println(arr)

	arr2 := make([]int, 5, 10)
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3

	_ = append(arr2, 4, 5) //append始终会返回一个新的切片，所以老的切片不管长度是否足够都不会改变
	fmt.Println(arr2)
}

func appendElement(arr []int, targets ...int) []int {
	for _, v := range targets {
		arr = append(arr, v)
	}
	return arr
}
