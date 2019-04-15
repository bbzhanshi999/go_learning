package main

import "fmt"

type number struct {
	f float32
}

type nr number

func main() {
	a := number{10.9}
	b := nr{20.0}
	c := number(b)

	fmt.Println(a, b, c)
}
