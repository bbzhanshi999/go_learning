package main

import "fmt"

func main() {
	slice1 := make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for i, v := range slice1 {
		fmt.Printf("index %d value is %d \n", i, v)
	}
}
