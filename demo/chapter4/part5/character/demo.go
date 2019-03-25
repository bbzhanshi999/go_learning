package main

import (
	"fmt"
	"unicode"
)

func main() {
	var ch int32 = '\u0041'
	var ch2 int32 = '\u03B2'
	var ch3 int32 = '\u7684'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	fmt.Printf("is letter %v", unicode.IsLetter('a'))
	fmt.Printf("is letter %v", unicode.IsLetter(rune(ch)))
}
