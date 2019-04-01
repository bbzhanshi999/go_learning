package main

import "fmt"

var sl []byte = []byte{'y', 'z'}

func main() {
	byteArr := []byte{'a', 'b', 'c'}

	fmt.Printf("%s", Append(sl, byteArr))
	
}

func Append(slice []byte, data []byte) []byte {
	s := make([]byte, len(slice)+len(data))

	for i := 0; i < len(slice); i++ {
		s[i] = slice[i]
	}

	for i := len(slice); i < len(s); i++ {
		s[i] = data[i-len(slice)]
	}

	return s
}
