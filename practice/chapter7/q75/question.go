package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}

	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)

	double(&items)
	fmt.Println(items)
}

func double(source *[5]int) {
	for i := 0; i < len(source); i++ {
		source[i] = source[i] * 2
	}

}
