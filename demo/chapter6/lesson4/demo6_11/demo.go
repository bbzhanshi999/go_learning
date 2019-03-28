package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	target()
	fun1("asd")
	fun2("hehehe")
}

func tracing(s string) string {
	fmt.Printf("start tracing func %s\n", s)
	return s
}

func un(s string) {
	fmt.Printf("leaving the func %s\n", s)
}

func target() {
	defer un(tracing("target")) //在这里，为什么tracing（）函数仍然能先执行，
	//还是由于defer中表达式引用的参数是在运行defer钱就记录了状态而已 等同与 methodName:=tracing("target"); defer un(methodName)
	fmt.Printf("running the target func\n")
}

func fun1(s string) (int, error) {

	n := 7
	err := io.EOF
	defer func() {
		log.Printf("fun1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF

}

func fun2(s string) (n int, err error) {
	//defer log.Printf("func1(%q) = %d, %v", s, n, err) //这种方式当即就获得了变量的状态，所以返回的n为0,空值

	//用返回匿名函数的方式和直接执行的效果完全不一样，类似与javascript的闭包，返回的是函数而不是函数结果，函数结果在运行完return后在执行
	//闭包方式不会立即执行函数，而是在return后执行，此时的nwei7,err不为空
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()

	return 7, io.EOF
}
