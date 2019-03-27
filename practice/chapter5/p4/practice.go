package main

import (
	"fmt"
	"strings"
)

func main() {
	practice5_4()
	practice5_5()
}

func practice5_4() {
	var i int
	for i = 0; i < 15; i++ {
	}
	fmt.Printf("index is %v", i)

	i = 0
BEGIN:
	i++
	if i < 15 {
		goto BEGIN
	}

	fmt.Printf("index is %v", i)
}

func practice5_5() {
	for i := 1; i <= 25; i++ {
		for j := 1; j < i; j++ {
			print("G")
		}
		println("")
	}

	for i := 1; i <= 25; i++ {
		println(strings.Repeat("A", i))
	}
}
