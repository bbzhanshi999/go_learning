package main

import "fmt"
import "./matrix"

func main() {
	fmt.Println(matrix.NewMatrix(1))
	// fmt.Println(new(matrix.matrix)) //编译失败
}
