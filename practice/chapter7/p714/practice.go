package main

import "fmt"

/*编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 `[]byte` 类型的切片）。

如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。

如果您想要反转 Unicode 编码的字符串，请使用 `[]int32` 类型的切片。*/
func main() {
	str := "Google"
	fmt.Println(reverse1(str))

	str = "我爱你"
	fmt.Println(reverse1(str))

}

func reverse1(source string) string {
	arr := []int32(source)

	j := 0

	for i := len(arr) - 1; i >= len(arr)/2; i-- {
		arr[i], arr[j] = arr[j], arr[i]
		j++
	}

	return string(arr)

}
