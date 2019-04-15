package main

import "fmt"
import "./even"

//*练习 9.4** 创建一个程序 main_oddven.go 判断前 100 个整数是不是偶数，将判断所用的函数编写在 even 包里。
func main() {
	for i := 0; i < 100; i++ {
		if even.IsEven(i) {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
