package main

import "fmt"

func main() {
	a := "haha"
	fmt.Printf("a  is %v\n", a)
	DoSomething(a)
	fmt.Printf("a  is %v\n", a)
	DoSomething1(&a)
	fmt.Printf("a  is %v\n", a)
}

func DoSomething(a string) {
	a = "hehe"
}

func DoSomething1(a *string) {
	*a = "heyhey"
}
