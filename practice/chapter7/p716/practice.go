package main

import "fmt"

//编写一个程序，使用冒泡排序的方法排序一个包含整数的切片（算法的定义可参考 [维基百科](http://en.wikipedia.org/wiki/Bubble_sort)）。
func main() {
	arr := []int{4, 40, 10, 50, 3, 2, 4, 7, 9, 8, 100}
	fmt.Println(arr)
	sortByBubble(arr)
	fmt.Println(arr)
}

func sortByBubble(source []int) {

	for i := 1; i < len(source); i++ {
		for j := 0; j < len(source)-i; j++ {
			if source[j] >= source[j+1] {
				source[j+1], source[j] = source[j], source[j+1]
			}
		}
	}
}
