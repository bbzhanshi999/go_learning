package main

func main() {
	subPrint(10)
}

func subPrint(val int) {
	if val >= 1 {
		println(val)
		subPrint(val - 1)
	}
}
