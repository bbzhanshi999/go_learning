package main

import "fmt"

//创建一个上面 `Car` 和 `Engine` 可运行的例子，并且给 `Car` 类型一个 `wheelCount` 字段和一个 `numberOfWheels()` 方法。

// 创建一个 `Mercedes` 类型，它内嵌 `Car`，并新建 `Mercedes` 的一个实例，然后调用它的方法。

// 然后仅在 `Mercedes` 类型上创建方法 `sayHiToMerkel()` 并调用它。

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (car *Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("hay ,Merkel")
}

func main() {
	m := Mercedes{Car{nil, 4}}
	m.Start()
	m.Stop()

	fmt.Println(m.numberOfWheels())

	m.sayHiToMerkel()
}
