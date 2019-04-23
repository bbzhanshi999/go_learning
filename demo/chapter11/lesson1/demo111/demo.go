package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

type Rectangle struct {
	width, height float32
}

func (re *Rectangle) Area() float32 {
	return re.width * re.height
}

func (re *Rectangle) Perimeter() float32 {
	return re.width*2 + re.height*2
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaInf Shaper
	areaInf = sq1
	fmt.Printf("the square has area: %f\n", areaInf.Area())
	fmt.Printf("the Perimeter has area: %f\n", areaInf.Perimeter())

	sqArr := []Shaper{sq1, &Rectangle{5.0, 6.0}}

	for _, v := range sqArr {
		fmt.Printf("Shape details: %v\n", v)
		fmt.Printf("Shape area is : %v\n", v.Area())
		fmt.Printf("Shape perimeter is : %v\n", v.Perimeter())
	}
}
