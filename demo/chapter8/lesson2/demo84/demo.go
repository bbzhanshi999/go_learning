package main

import "fmt"

// 演示检查key是否存在以及删除某个key
func main() {
	map1 := make(map[string]int)
	map1["Buffalo"] = 1
	map1["Ty Gwyn"] = 2
	map1["Longdon"] = 3

	if v, ok := map1["Ty Gwyn"]; ok {
		fmt.Printf("%d is Bily's hometown\n", v)
	}

	if _, ok := map1["Washington"]; !ok {
		fmt.Println("there is no Washington")
	}

	delete(map1, "Longdon")
	if _, ok := map1["Longdon"]; !ok {
		fmt.Println("there is no Longdon")
	}
}
