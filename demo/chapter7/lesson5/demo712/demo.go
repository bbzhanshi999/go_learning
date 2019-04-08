package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	sl_to[0] = 100
	sl_to[1] = 100
	sl_to[2] = 100
	sl_to[3] = 100
	sl_to[4] = 100
	count := copy(sl_to, sl_from)
	fmt.Println(count)
	fmt.Println(sl_to)

	sl3 := []int{1, 3, 4}
	sl3 = append(sl3, 6, 7, 8)
	fmt.Println(sl3)
}
