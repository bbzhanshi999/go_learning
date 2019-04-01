package main

import "fmt"

func main() {
	s1, s2 := cutBuf([]byte{'a', 'b', 'c', 'd', 'e'}, 2)
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)
}

func cutBuf(buf []byte, n int) ([]byte, []byte) {
	return buf[:n], buf[n:]
}
