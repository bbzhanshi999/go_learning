package main

import "fmt"

//假设有字符串 str，那么 `str[len(str)/2:] + str[:len(str)/2]` 的结果是什么？
func main() {
	str := "123456789"
	fmt.Println(str[len(str)/2:] + str[:len(str)/2])
	//相当与从中间切分后在和起来
}
