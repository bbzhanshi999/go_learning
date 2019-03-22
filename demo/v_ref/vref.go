package main

import "fmt"

func main() {
	var a int
	a = 1
	c := &a
	d := &a
	fmt.Printf("a: %p b:%p", c, d)
	fmt.Printf("a: %v b:%v", *c, *d)
	fmt.Printf("a: %v", &c)
}
