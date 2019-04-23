package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s *stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c *car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(v *valuable) {
	fmt.Printf("the value is %f\n", (*v).getValue())
}

func main() {
	o := stockPosition{"goog", 577.30, 4}
	var i valuable
	i = &o
	showValue(&i)

	i = &car{"BMW", "M3", 66500}
	showValue(&i)

}
