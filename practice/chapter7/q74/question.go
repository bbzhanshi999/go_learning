package main

import "fmt"

//问题 7.4** 假设 `s1 := []byte{'p', 'o', 'e', 'm'}` 且 `s2 := s1[2:]`，s2 的值是多少？如果我们执行 `s2[1] = 't'`，s1 和 s2 现在的值又分别是多少？
func main() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s2)

	fmt.Println(s1)

	s2[1] = 't'

	fmt.Println(s2)
	fmt.Println(s1)
}
