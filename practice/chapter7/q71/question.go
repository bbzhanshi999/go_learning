package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d", "e"} //...表示根据实际传入的参数设置length
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	println(len(a))
	//a[5] = "f" //由于长度是五，这地方会报数组角标越界
}
