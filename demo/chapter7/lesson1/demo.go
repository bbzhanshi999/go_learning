package main

import "fmt"

func main() {
	pointerOrValue()

	a := [...]int{1, 2, 3, 4, 5}

	printValue(a)
	fmt.Println(a) //改变的是a的一个复制，因此不会动影响原来的数据
	printPointer(&a)
	fmt.Println(a) //改变是的a的指针对应的值，因此会影响原来的值

}

//演示创建数组和数组指针的不同
func pointerOrValue() {
	//创建数组
	var arr1 [5]int

	//创建数组指针
	arr2 := new([5]int)

	arr2[0] = 1
	arr2[1] = 2

	// arr1 = arr2 // cannot use arr2 (type *[5]int) as type [5]int in assignment

	//将arr2的值的拷贝复制给了arr1
	arr1 = *arr2

	println(arr1[0])

	arr2[0] = 10

	println(arr1[0])
	println(arr2[0])

}

func printValue(a [5]int) {
	a[4] = 100
	fmt.Println(a)
}
func printPointer(a *[5]int) {
	a[4] = 100
	fmt.Println(a)
}


