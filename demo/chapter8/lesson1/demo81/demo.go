package main

import "fmt"

//map创建的几种方式
func main() {

	var mapLit map[string]int
	//var mapCreated  map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2} //声明式的创建方式 ，类似与json
	mapCreated := make(map[string]float32)      //通过make方式创建，注意和切片的区别，不需要加入长度和容量参数
	mapAssigned = mapLit                        //通过分配指针的方式赋值，此时mapAssigned和mapLit共享内存

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14
	mapAssigned["third"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["third"]) //虽然没有给mapLit的third赋值，但由于和mapAssigned共享内存
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])    //如果没有ten，那么go会给该key附上默认零值

}
