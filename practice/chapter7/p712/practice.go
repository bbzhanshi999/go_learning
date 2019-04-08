package main

import "fmt"

//写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。
func main() {
	origin := []string{"h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"}
	fmt.Println(RemoveStringSlice(2, 5, origin))
}

func RemoveStringSlice(start int, end int, origin []string) []string {

	result := make([]string, len(origin)-(end-start+1))

	count := copy(result, origin[:start])
	copy(result[count:], origin[end+1:]) //起步的时候不能包含end本身的字符串元素
	return result
}
