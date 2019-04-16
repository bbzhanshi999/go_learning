package main

import (
	"fmt"
	"unsafe"
)

type File struct {
	desc int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	file := NewFile(1, "1.jpg")
	fmt.Println(unsafe.Sizeof(*file)) //实例本身的内存大小
	fmt.Println(unsafe.Sizeof(file))  //指针的大小
}
