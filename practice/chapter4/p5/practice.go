package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"

	fmt.Printf("str1 byte length: %v \n", len(str1))
	fmt.Printf("str1 string length: %v \n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 byte length: %v \n", len(str2))
	fmt.Printf("str2 string length: %v \n", utf8.RuneCountInString(str2))

	//通过本练习熟悉字节与字符的区别，在UTF8编码中，一个字符的字节数是不定长的（1-4字节，每个字节二进制长度为8位），
	//字节长度用native函数len（），字符长度需要根据编码调用unicode包的RuneCountInString进行字符长度计算
}
