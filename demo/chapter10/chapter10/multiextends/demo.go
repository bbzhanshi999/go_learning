package main

import "fmt"

type Camera struct {
}

func (camera *Camera) TakePicture() string {
	return "click"
}

type Phone struct {
}

func (phone *Phone) Call() string {
	return "ring,ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println(cp.Call())
	fmt.Println(cp.TakePicture())
}
