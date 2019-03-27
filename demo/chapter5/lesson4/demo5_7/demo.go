package main

import "fmt"

func main() {
	str := "this is an English sentence"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str's %vth word is %c\n", i, str[i])
	}

	str2 := "四个汉字"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("str2's %vth word is %c\n", i, str2[i]) //由于ascII码只占用一个字节，所以要迭代字符串必须要用其他方式
	}
}
