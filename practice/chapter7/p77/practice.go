package main

import "fmt"

/*
a) 写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，返回该数组的所有数字和。

如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？

b) 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。
*/
func main() {

	arrF := [...]float32{1.0, 2.0, 3.0}
	rs := Sum(arrF[:len(arrF)])
	fmt.Println(rs)

	arr := []int{1, 2, 3, 4, 5, 6}
	ave, sum := SumAndAve(arr)
	fmt.Printf("average is %f, sum is %d", ave, sum)
}

func Sum(arrF []float32) (rs float32) {
	for _, v := range arrF {
		rs += v
	}
	return
}

func SumAndAve(arr []int) (ave float32, sum int) {
	for _, v := range arr {
		sum += v
	}
	ave = float32(sum / len(arr))
	return
}
