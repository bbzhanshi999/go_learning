package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "this is a part of content that contains 中文"

	fmt.Printf("the length of byte %d \n", len(str))
	fmt.Printf("the leagth of character %d \n", utf8.RuneCountInString(str))
	fmt.Printf("the leagth of character %d \n", len([]int32(str)))

	fmt.Printf("%c", []rune(str)[41]) //如果要找到字符串中某一个具体位置的字符，现将str转换成 rune数组，然后取指定位置的字节
	var b []byte
	b = append(b, str...) //str...相当于将str转换成byte数组并且作为可变参数传进去

	fmt.Println(b)

	c := []rune{'我', '爱', '你'}
	c = append(c, c...) //c...的含义为将数组转换以可变参数的形式传入方法，前提是方法可以接受可变参数
	fmt.Printf("%c\n", c)

	fmt.Printf("sub string from 0 to 10 is \"%s\" \n", str[0:10]) //通过类似切片拆分的语法来对字符串进行截取

	//修改字符串。
	s := `hello`

	sArr := []rune(s)

	sArr[0] = '哈'

	s = string(sArr)

	fmt.Println(s)

	s1 := "我日"
	s2 := "我日"

	fmt.Printf("s1 equal s2 %v\n", s1 == s2)

}
