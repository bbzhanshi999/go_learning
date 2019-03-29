package main

import "fmt"

func main() {
	f()
	a := f2() //a 等于2 ，因为defer在return后执行，由于闭包函数内部拿到的ret是在defer后函数执行时才去拿值
	println(a)
}

func f() {
	for i := 0; i < 4; i++ {
		g := func() {
			fmt.Printf("%d ", i)
		}
		g()
		fmt.Printf("g  s of type %T and has value %v \n", g, g)
	}

}

func f2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
