package main

import (
	"fmt"
	"os"
	"strconv"
)

const Pi = 3.1415926

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	RED Color = iota
	ORANGE
)

var inta, intb int

var (
	intv   int
	boolv  bool
	floatv float32
	strv   string
	userv  User
)

type Color int

type User struct {
	username string
	password string
}

func main() {
	fmt.Printf("%v \n", Pi*3)
	gender := 2
	switch gender {
	case Unknown:
		fmt.Println("未知")
	case Female:
		fmt.Println("女")
	case Male:
		fmt.Println("男")
	}

	fmt.Printf("value: %v \n", strconv.FormatInt(2, 10))

	fmt.Printf("a: %v b: %v c: %v \n", a, b, c)

	fmt.Printf("orange is : %v \n", ORANGE)

	fmt.Printf("%v %v %v \n", inta, intb, inta+10)

	fmt.Printf("nil value of each type: %v %v %v %v %v \n", intv, floatv, boolv, strv, userv)

	goPath := os.Getenv("GOROOT")

	fmt.Println(goPath)

}
