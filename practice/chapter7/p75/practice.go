package main

var sl []byte = []byte{'b', 'b'}

func main() {
	byteArr := []byte{'a', 'b', 'c'}

	Append(sl, byteArr)

}

func Append(slice []byte, data []byte) []byte {
	s := make([]byte, len(slice)+len(data))
	// todo
	return s
}
