package main

import "fmt"

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	i int
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(i int) {
	s.i = i
}

func callMethod(s Simpler) {
	fmt.Printf("simpler get value is: %d\n", s.Get())
	s.Set(10)
}

func main() {
	s := Simple{i: 20}

	var si Simpler
	si = &s
	callMethod(si)

	fmt.Println(si.Get()) //对于go而言， 接口的至实际上是一个保存了实例指针的内存空间，因此即使在callMethod的方法中我们调用了接口的拷贝的set方法，但是作用的仍然是实例指针指向的那个实例
}
