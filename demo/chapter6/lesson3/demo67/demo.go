package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("min value is %v\n", x)
	slice := []int{7, 6, 7, 2, 8}
	x = min(slice...) //可变参数可以直接传一个slice进去
	fmt.Printf("min value is %v\n", x)
}

func min(s ...int) int {
	min := s[0]
	for i := 0; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
	}
	return min
}
