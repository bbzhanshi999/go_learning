package main

import "fmt"

//给定一个slice`s []int` 和一个 int 类型的因子，扩展 s 使其长度为 `len(s) * factor`。
func main() {
	arr := []int{1, 2, 3}
	fmt.Println(expandLen(arr, 10))
}

func expandLen(arr []int, factor int) []int {
	newSlice := make([]int, len(arr)*factor)
	copy(newSlice, arr)
	return newSlice
}
