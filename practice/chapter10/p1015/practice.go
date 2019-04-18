package main

import "fmt"

//为 int 定义别名类型 `TZ`，定义一些常量表示时区，比如 UTC，定义一个 map，
//它将时区的缩写映射为它的全称，比如：`UTC -> "Universal Greenwich time"`。为类型 `TZ` 定义 `String()` 方法，它输出时区的全称。

type TZ int

const (
	UTC TZ = 1
)

var tzMap map[TZ]string = map[TZ]string{UTC: "Universal Greenwich time"}

func (t TZ) String() string {
	return tzMap[t]
}

func main() {
	fmt.Println(UTC)
}
