package main

func main() {
	f1()

	i := 0
	defer println(i)
	i++
	println(i)
	f()

	//defer 后的表达式如果要使用参数，那么规则是记录当前的参数状态，最后执行的时候取出当时的这个状态
	//defer按照先进后出的顺序来完成代码的执行
}

func f1() {
	println("function 1 first process")
	defer f2()
	println("function 1 stopped")
}

func f2() {
	println("function 2 process")
}

func f() {
	for i := 0; i < 4; i++ {
		defer println(i)
	}
}
