package main

import "strings"

func main() {
	println(strings.Map(replace, "我is your 爸爸"))
}

func replace(r rune) rune {
	if !IsAsc(r) {
		return rune('?')
	}
	return r
}

func IsAsc(c rune) bool {
	if c < 255 {
		return true
	}
	return false
}
