package main

import (
	"fmt"
	"strings"
)

func main() {
	practice5_4()
	// practice5_5()
	//practice5_6()
	//practice5_7()
	practice5_8(20, 10)
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

func practice5_6() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v the complement of %b is: %b\n", i, i, ^i)
	}
}

func practice5_7() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			println("fizzbuzz")
		case i%5 == 0:
			println("buzz")
		case i%3 == 0:
			println("fizz")
		default:
			println(i)
		}
	}
}

// 打印 rectangle
func practice5_8(width int, height int) {
	for i := 0; i < height; i++ {
		println(strings.Repeat("*", width))
	}
}
