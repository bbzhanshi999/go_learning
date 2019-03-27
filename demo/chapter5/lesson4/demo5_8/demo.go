package main

import "fmt"

func main() {
	i := 5
	for i >= 0 {
		i -= 1
		fmt.Printf("the variable i is now:%d\n", i)
	}
	infinite() //无限循环

}

func infinite() {
	for {
		print(".")
	}
}

//无限循环经典应用
// for t, err = p.Token(); err == nil; t, err = p.Token() {
// 	...
// }
