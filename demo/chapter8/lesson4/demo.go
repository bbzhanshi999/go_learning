package main

import "fmt"

//假设我们想获取一个 map 类型的切片，我们必须使用两次 `make()` 函数，第一次分配切片，第二次分配 切片中每个 map 元素（参见下面的例子 8.4）。
func main() {
	items := make([]map[int]int, 5)

	//要想给每个元素赋值，不能用第二种方式
	for i := range items {
		items[i] = make(map[int]int)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	//这种方法获得的item只是一个拷贝，因此赋值是无效的
	items2 := make([]map[int]int, 5)

	for _, item := range items2 {
		item = make(map[int]int)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2) //打印结果都为空值，因为并没有获得没一个元素的指针，而是拷贝
}
