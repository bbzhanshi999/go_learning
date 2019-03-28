package main

import (
	"strings"
	"unicode"
)

func main() {
	callback(1, add)
	println(strings.IndexFunc("我日ab", IsAsc))
	println(strings.IndexFunc("asd12", unicode.IsNumber))
}

func add(a, b int) {
	println(a + b)
}

func callback(a int, f func(int, int)) {
	f(a, 2)
}

func IsAsc(c rune) bool {
	if c < 255 {
		return true
	}
	return false
}
