package main

import "fmt"

//编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
func main() {
	str := "我爱你，亲爱的姑娘"
	front, end := splitString(str, 3)
	fmt.Printf("%s    %s", front, end)
}

func splitString(source string, i int) (front string, end string) {
	runArr := []rune(source)
	return string(runArr[:i]), string(runArr[i+1:])
}
