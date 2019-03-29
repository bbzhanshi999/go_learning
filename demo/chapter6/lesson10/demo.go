package main

import (
	"log"

	"./hello"
)

func main() {
	// where := func() {
	// 	counter, file, line, ok := runtime.Caller(1) //Caller的入参：0 代表当前的方法栈，1,代表地一个父栈 一次类推
	// 	log.Printf("%d:%s:%d:%v", counter, file, line, ok)
	// }
	log.SetFlags(log.Llongfile)

	hello.Hello()
	//where() //通过闭包的方式可以在闭包函数内获得对因当前的
	log.Print("\n")
	hello.Hi()
	//where()
	log.Print("\n")

}
