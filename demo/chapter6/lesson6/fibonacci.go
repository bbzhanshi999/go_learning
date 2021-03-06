package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 41; i++ {
		idx, result := fibonacci(i)
		fmt.Printf("index %v value is %v \n", idx, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("fibonacci process %s", delta)
}

func fibonacci(idx int) (i int, res int) {
	if idx <= 1 {
		return idx, 1
	}
	_, p1 := fibonacci(idx - 1)
	_, p2 := fibonacci(idx - 2)
	return idx, p1 + p2

}
