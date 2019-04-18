package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
}

type D struct {
	B
	b float32
}

func main() {
	outer := new(outerS)
	outer.b = 10
	outer.c = 1.11
	outer.int = 10
	outer.in1 = 1
	outer.in2 = 2

	fmt.Println(outer.innerS)

	outer.innerS = *new(innerS)
	outer.innerS.in1 = 10
	outer.innerS.in2 = 20
	fmt.Println(outer.innerS)

	c := new(C)
	c.A.a = 1
	c.B.a = 2
	fmt.Println(c)

	d := new(D)
	d.B.b = 10
	d.a = 10
	d.b = 10.1
	fmt.Println(d)
}
