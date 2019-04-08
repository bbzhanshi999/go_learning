package main

import "fmt"

//编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
func main() {
	str := "aaaaabbbbbbbccccccc"
	fmt.Println(string(findDefer([]int32(str))))
}

func findDefer(source []int32) []int32 {
	var target []int32
	for i, v := range source {
		if i > 0 {
			if v != source[i-1] {
				target = append(target, v)
			}
		}
	}
	return target
}
