package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	mapFunc(func(i int) int {
		return i * 100
	}, arr)
	fmt.Println(arr)
}

func mapFunc(fn func(int) int, source []int) {
	for i := 0; i < len(source); i++ {
		source[i] = fn(source[i])
	}
}
