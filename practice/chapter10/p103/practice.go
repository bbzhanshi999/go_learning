package main

import (
	"fmt"
	"math"
)

//使用坐标 X、Y 定义一个二维 Point 结构体。同样地，对一个三维点使用它的极坐标定义一个 Polar 结构体。
//实现一个 `Abs()` 方法来计算一个 Point 表示的向量的长度，实现一个 `Scale` 方法，它将点的坐标乘以一个尺度因子
type Point struct {
	x, y float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func Scale(p Point, s float64) Point {
	p.x = p.x * s
	p.y = p.y * s
	return p
}

func main() {
	p := new(Point)
	p.x = 1.0
	p.y = 2.0
	fmt.Println(Abs(p))

	fmt.Println(Scale(*p, 5)) //在这里，传入的是p的一个拷贝，因此不会对元数据造成修改
	fmt.Println(p)
}
