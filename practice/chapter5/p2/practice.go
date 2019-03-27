package main

import (
	"log"
	"math"
)

//  实现开根号
func main() {
	if v, ok := mySqrt(-9); ok {
		println(v)
	} else {
		log.Printf("has error")
	}
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return //这里有由于返回为空，所以go会自动将对应类型的空值返回，bool默认为false，数字类型为0
	}
	return math.Sqrt(f), true
}
