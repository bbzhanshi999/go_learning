package main

import (
	"container/list"
	"fmt"
)

//不能给非本包的类型定义方法
// func (p *list.List) Iter() {
// 	// ...
// }

//但是有一个间接的方式：可以先定义该类型（比如：int 或 float）的别名类型，
//然后再为别名类型定义方法。或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
type myList list.List

func (p *myList) Iter() {
	//...
}

func main() {
	lst := new(myList)
	for _ = range lst.Iter() {
		fmt.Println("")
	}
}
