package main

import (
	"fmt"
	"strconv"
)

type Celsius float64

//String方法必须用类型作为接受者
func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 2, 64) + "°C"
}

func main() {
	c := Celsius(64.3)
	fmt.Println(c)
}
