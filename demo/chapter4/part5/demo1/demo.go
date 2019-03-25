package main

import "fmt"

func main() {
	var avar = 10
	println(avar == 5)
	println(avar == 10)
	println(avar == 5 && avar == 10)
	println(avar == 5 || avar == 10)

	fmt.Printf("boolean value is :%t", avar == 5)

	var inta int32 = 666666
	println(inta)

}
