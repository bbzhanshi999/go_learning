package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func Area(r *Rectangle) int {
	return r.height * r.width
}

func Perimeter(r *Rectangle) int {
	return r.height*2 + r.width*2
}

func main() {
	r := &Rectangle{4, 5}
	fmt.Println(Area(r))
	fmt.Println(Perimeter(r))
}
