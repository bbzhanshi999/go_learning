package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	exitWhenError()
	println("has not error") //由于已经推出进程，所以该行不会执行
}

func exitWhenError() {
	result, err := strconv.Atoi("hahaha")
	if err != nil {
		fmt.Printf("some error has occured :%v", err)
		os.Exit(1) //手动退出当前进程，并且设置状态码为1
	} else {
		fmt.Printf("%v", result)
	}
}
