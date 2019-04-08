package main

import "fmt"

//写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
func main() {
	origin := []int{1, 2, 3, 4, 5, 6, 7}
	target := []int{10, 20, 30, 40}
	fmt.Println(InsertStringSlice(origin, target, 3))
}

func InsertStringSlice(origin []int, target []int, position int) []int {

	//1 先将元切片按照位置拆分成两个切片
	front, back := origin[:position], origin[position:]
	//2 用append方法插入target和后续的元素
	var result []int = make([]int, len(origin)+len(target))

	copy(result, front)	
	copy(result[len(front):], target)
	copy(result[len(front)+len(target):], back)

	return result
}
