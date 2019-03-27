package main

import "fmt"
import "./season"

func main() {
	test1()
	result, err := season.Season(11)
	if err != nil {
		fmt.Printf("%e", err)
	} else {
		println(result)
	}

}

func test1() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
