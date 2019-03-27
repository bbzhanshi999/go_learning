package main

import "fmt"

func main() {
	//通过for range 进行迭代可以获得字符而不是字节
	str := "Go is a beautiful language!"
	for i, char := range str {
		fmt.Printf("str %vth word is %c\n", i, char)
	}

	str2 := "Chinese: 日本語"
	for i, char := range str2 {
		fmt.Printf("str2 %vth word is %c\n", i, char)
		fmt.Printf("%d      %d      %U '%c' %X\n", i, char, char, char, []byte(string(char)))
	}

	// s := ""
	// for s != "aaaaa" {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }

	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
	// 	s = i+1, j+1, s+"a" {
	// 	fmt.Println("Value of i, j, s:", i, j, s)
	// }
}
