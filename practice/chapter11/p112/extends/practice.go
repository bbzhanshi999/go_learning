package main

import (
	"fmt"
	"math"
)

//b) 使用一个抽象类型 `Shape`（没有字段） 实现同样的功能，它实现接口 `Shaper`，然后在其他类型里内嵌此类型。扩展 10.6.5 中的例子来说明覆写。

type Shaper interface {
	Area() float32
	Perimeter() float32
}

type Shape struct{}

func (s Shape) Area() float32 {
	return -1
}

func (s Shape) Perimeter() float32 {
	return -1
}

type Square struct {
	side float32
	Shape
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

type Rectangle struct {
	width, height float32
	Shape
}

func (re *Rectangle) Area() float32 {
	return re.width * re.height
}

func (re *Rectangle) Perimeter() float32 {
	return re.width*2 + re.height*2
}

type Circle struct {
	radius float32
	Shape
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	s := Shape{}
	sq := Square{10, s}
	r := Rectangle{10, 20, s}
	c := Circle{10, s}
	shapes := []Shaper{&sq, &r, &c}
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
		fmt.Println("Perimeter of this shape is: ", shapes[n].Perimeter())
	}
}
