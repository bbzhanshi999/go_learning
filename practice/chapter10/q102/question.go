package main

import "fmt"

// a）假设定义： `type Integer int`，完成 `get()` 方法的方法体: `func (p Integer) get() int { ... }`。

// b）定义： `func f(i int) {}; var v Integer` ，如何就 v 作为参数调用f？

// c）假设 `Integer` 定义为 `type Integer struct {n int}`，完成 `get()` 方法的方法体：`func (p Integer) get() int { ... }`。

// d）对于新定义的 `Integer`，和 b）中同样的问题。

type Integer int

func (i Integer) get() int {
	return int(i)
}

func f(i int) {
	fmt.Println(i)
}

var v Integer

type Integer2 struct {
	n int
}

func (i Integer2) get() int {
	return i.n
}

var v2 Integer2

func main() {
	v = 10
	v2.n = 20
	f(int(v))
	f(v2.get())
}
