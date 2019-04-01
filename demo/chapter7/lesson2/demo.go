package main

import (
	"bytes"
	"fmt"
)

var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {

	var slice = arr[0:5]
	printlnSlice(slice)

	//println(slice[5])
	println(arr[5])
}

func printlnSlice(slice []int) {
	// println(len(slice))
	// println(cap(slice))
	// println(&slice)
	// fmt.Println(reflect.TypeOf(slice))

	// fmt.Println(reflect.TypeOf(arr))

	// slice = slice[:cap(slice)] //扩大切片的大小上限，最大只能为数组的长度，要注意slice的指针发生了变化
	// slice[5] = 999
	// println(slice[5])

	bufferDemo()
}

func bufferDemo() {
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Println(buffer.String())
}

var str = "hello go,我日你妈"
var flag = 0

func getNextString() (string, bool) {
	r := []rune(str) //要想去的string的制定index字符，并且再转换成string，就要先将string转换为[]rune切片，再将指定切片通过string（）构造成字符串
	if flag < len(r) {
		flag++
		return string(r[flag-1 : flag]), true
	}
	return "", false
}
