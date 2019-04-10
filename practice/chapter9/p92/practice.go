package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 129
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(99999999999999))
}
