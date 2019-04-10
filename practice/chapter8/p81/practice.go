package main

import "fmt"

//创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday。
func main() {
	days := map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday ", 7: "Sunday"}

	for k, v := range days {
		fmt.Println(k, "is", v)
	}

	if _, ok := days[2]; ok {
		fmt.Println("have 2")
	}

}
