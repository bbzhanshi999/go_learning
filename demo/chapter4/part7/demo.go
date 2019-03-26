package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串strings工具
	stringsTest()
	//字符串类型转换
	stringconvTest()
}

func stringsTest() {
	str1 := "四个汉字"
	println(strings.HasPrefix(str1, "四"))
	println(strings.HasSuffix(str1, "字"))
	println(strings.Contains(str1, "个"))

	str2 := "four chinese word,四个汉字"
	println(strings.Index(str2, "四"))
	println(strings.IndexRune(str2, rune('四'))) //查询非ascII码字符
	println(strings.LastIndex(str2, "o"))

	str3 := "goooooogle"
	println(strings.Replace(str3, "o", "a", 4))  //gaaaaoogle
	println(strings.Replace(str3, "o", "a", -1)) //gaaaaaagle -1代表替换所有
	println(strings.ReplaceAll(str3, "o", "a"))  //gaaaaaagle
	println(strings.Count(str3, "o"))            //6

	str4 := "    gooooogle     "
	println(strings.TrimSpace(str4))
	println(strings.Trim(str3, "g")) //修剪指定开头和结尾的字符串
	strs := strings.Split(str4, "o") //split方法切分成了切片
	for result := range strs {
		println(strs[result])
	}

	reader := strings.NewReader(str1)
	ru, size, _ := reader.ReadRune() //读取一个字符
	println(ru)
	println(size)
	ru, size, _ = reader.ReadRune()
	println(ru)
	println(size) //size 代表的是字节长度
}

var stringconvTest = func() {
	println(strconv.Itoa(10))
	println(strconv.FormatFloat(110.123232, 'b', 2, 32))
	println(strconv.FormatFloat(110.123232, 'b', 4, 64))
	println(strconv.FormatFloat(11.123232, 'e', 3, 64))
	println(strconv.FormatFloat(110.123232, 'e', 2, 64))
	println(strconv.FormatFloat(110.123232, 'f', 2, 64))
	println(strconv.FormatFloat(110.123232, 'g', 2, 64))

	i, err := strconv.Atoi("111")
	fmt.Printf("%v\n,%v\n", i, err)
	ep := &err
	f, err := strconv.ParseFloat("A5435.46546237678", 32)
	fmt.Printf("%v\n,%v\n", f, err)
	fmt.Printf("%v", *ep == err)
	fmt.Printf("%v", &ep)
}
