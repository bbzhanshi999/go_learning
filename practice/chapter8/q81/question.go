package main

import "fmt"

//问题 8.1： 下面这段代码的输出是什么？
func main() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		//fmt.Printf("%s's capital is %s\n", key, capitals[key])
		fmt.Println("Capital of", key, "is", capitals[key]) //fmt.Println 可以多参数组合，并且用空格隔开每个参数
	}
}
