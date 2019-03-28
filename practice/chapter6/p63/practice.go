package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "aaaa"

	t := reflect.TypeOf(a)
	fmt.Printf("%v \n", t)

	typecheck("1", 1.1, 2, 'a')
}

func typecheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
		case float64:
		case string:
		case bool:
		default:
			fmt.Printf("%v \n", v)
		}
	}
}
