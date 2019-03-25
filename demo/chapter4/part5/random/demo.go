package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}

	for i := 0; i < 5; i++ {
		r := rand.Intn(8) //8以内的数字
		fmt.Printf("%d / ", r)
	}

	fmt.Println()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}
}
