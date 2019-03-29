package main

import (
	"fmt"
	"time"
)

var fibs [42]int

//通过缓存存储已经计算得到值比直接进行递归重新计算优势明显
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

	if fibs[idx] != 0 {
		return idx, fibs[idx]
	}

	if idx <= 1 {
		fibs[0] = 1
		return idx, 1
	}
	_, p1 := fibonacci(idx - 1)
	_, p2 := fibonacci(idx - 2)
	fibs[idx] = p1 + p2
	return idx, fibs[idx]

}
