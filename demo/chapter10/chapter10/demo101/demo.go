package main

import (
	"fmt"
	"reflect"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Interval struct {
	start int
	end   int
}

func main() {
	ms := new(struct1)
	ms.i1 = 1
	ms.f1 = 0.12
	ms.str = "haha"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
	fmt.Println(reflect.TypeOf(ms))

	//混合字面量语法（composite literal syntax）
	ms2 := struct1{1, 1.1, "heeh"}
	fmt.Println(reflect.TypeOf(ms2))

	ms3 := &struct1{2, 2.2, "huohuo"}
	fmt.Println(reflect.TypeOf(ms3))

	intr := Interval{0, 3}
	//intr := Interval{end:5, start:1}
	//intr := Interval{end:5}
	intr.end = 10
	fmt.Println(intr)
}
