package main

import "fmt"

//演示数组常量键值对赋值法
func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{2: i, 1: i * i, 0: i * i * i})
	}
}

func fp(a *[3]int) { fmt.Println(a) }
