package main

import (
	"fmt"
	"time"
)

//定义结构体 Address 和 VCard，后者包含一个人的名字、地址编号、出生日期和图像，试着选择正确的数据类型。构建一个自己的 vcard 并打印它的内容。

type Vcard struct {
	name      string
	addresses []*Address
	birthday  *time.Time
	image     string
}

type Address struct {
	code   string
	detail string
}

func main() {
	var card Vcard
	card.name = "zhao"
	birthday := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)
	card.birthday = &birthday
	card.image = "root/profiles/1.jpg"
	card.addresses = []*Address{&Address{"210048", "大厂"}, &Address{"210015", "雄周"}}

	fmt.Println(card)
	fmt.Println(card.birthday)
	fmt.Println(card.addresses)
}
