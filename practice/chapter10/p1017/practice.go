package main

import "fmt"
import "./stack"

func main() {
	s := new(stack.Stack)
	s.Push(10)
	//fmt.Println(s.Pop())

	s.Push(20)
	//fmt.Println(s.Pop())

	s.Push(30)
	//fmt.Println(s.Pop())

	s.Push(40)
	//fmt.Println(s.Pop())

	s.Push(50)
	//fmt.Println(s.Pop())

	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
}
