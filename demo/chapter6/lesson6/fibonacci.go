package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		idx, result := fibonacci(i)
		fmt.Printf("index %v value is %v \n", idx, result)
	}
}

func fibonacci(idx int) (i int, res int) {
	if idx <= 1 {
		return idx, 1
	}
	_, p1 := fibonacci(idx - 1)
	_, p2 := fibonacci(idx - 2)
	return idx, p1 + p2

}
