package main

import "fmt"

//写一个 minSlice 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice 方法返回最大值。
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(minSlice(arr))
	fmt.Println(maxSlice(arr))

	fmt.Println(arr[1:6]) // 切片重组最多只能到原切片的长度
}

func minSlice(arr []int) int {
	min := arr[0]

	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	return min
}

func maxSlice(arr []int) int {
	max := arr[0]

	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}
