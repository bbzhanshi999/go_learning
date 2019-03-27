package main

import (
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//在if语句开头声明的变量作用域只能在控制流语句代码块中
	if val := 10; val < 20 {
		println("value " + strconv.Itoa(val) + " is less than 20")
	}

	if value := process(); value < 20 {
		println("value " + strconv.Itoa(value) + " is less than 20")
	} else {
		println("value " + strconv.Itoa(value) + " is bigger than 20")
	}
}

func process() int {
	// 防止rand产生的随机数重复
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)

	a := rand.Intn(100)
	return a
}
