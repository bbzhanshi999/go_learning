package main

import (
	"fmt"
	"math"
)

func main() {
	// test1()
	// n, err := uint8FromInt(2147483647)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// }
	// println(n)

	// println(IntFromFloat64(4294967295.111))
	// test2()

	test3()
}

func test1() {
	var n int16 = 34
	var m int32
	m = int32(n)
	var o = 12.6
	fmt.Printf("32 bit int is %d\n", m)
	fmt.Printf("16 bit int is %d\n", n)
	fmt.Printf("dasd %5.2g\n", o)
	fmt.Printf("dasd %5.2e\n", o)
}

func uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of unit8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func test2() {
	var a = 7
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", 10)
	fmt.Printf("%b\n", ^10)
	fmt.Printf("%v\n", ^uint(10))
	fmt.Printf("%b\n", ^uint(10))
}

type ByteSize float64

const (
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type BitFlag int

const (
	Active  BitFlag = 1 << iota //1<<0 = 1
	Send                        // 1<<1 = 3
	Receive                     //1MM2 =4
)

func test3() {
	fmt.Printf("Active | Send : %v\n", Active|Send)
	fmt.Printf("Active is %b\n", Active)
	fmt.Printf("Send is %b\n", Send)
}
