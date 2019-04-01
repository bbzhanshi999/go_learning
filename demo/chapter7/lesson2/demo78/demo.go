package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 10, 20) // make创建分片，参数：类型，长度，容量 等同于  [20]int[0:10]

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 5
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d \n", i, slice1[i])
	}

	fmt.Printf("length is %d", len(slice1))
	fmt.Printf("capacity is %d", cap(slice1))
}
