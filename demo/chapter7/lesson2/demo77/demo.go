package main

import "fmt"

func main() {
	var arr [6]int
	var slice1 []int = arr[2:5]

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d \n", i, slice1[i])
	}

	fmt.Printf("arr length %d\n", len(arr))
	fmt.Printf("slice length %d\n", len(slice1))
	fmt.Printf("slie capacity %d\n", cap(slice1))

	slice1 = slice1[0:cap(slice1)] //增大slice1的长度。最大长度为slice的容量也就是arr的长度,切片只能向后移动
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d \n", i, slice1[i])
	}
	fmt.Printf("slice length %d\n", len(slice1))
	fmt.Printf("slie capacity %d\n", cap(slice1))

	//注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!
}
