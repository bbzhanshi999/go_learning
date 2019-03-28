package main

import "fmt"

func main() {
	printResult(threeResult1(19, 10))
	printResult(threeResult2(19, 10))
}

func printResult(add, mult, sub int) {
	fmt.Printf("add is %d,mult is %d, sub is %d\n", add, mult, sub)
}

func threeResult1(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func threeResult2(a, b int) (add int, mult int, sub int) {
	add = a + b
	mult = a * b
	sub = a - b
	return
}
