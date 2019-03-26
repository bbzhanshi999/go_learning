package main

import (
	"fmt"
	"strings"
)

func main() {
	println(len("四个汉字"))                 //len打印的是字节长度
	println("我A this is a fuck damn"[0]) //通过[]括号来获得字节，只能对ascII码字符有效，因为只有一个字节，而UTF-8的字节数在1-4个

	println(strings.Join([]string{"a", "b"}, ",")) //拼接字符
	str := fmt.Sprintf("%v", 1)                    //可以将格式化的值返回回来，可以利用此进行各种数据类型的转字符串工作
	println(str)
}
