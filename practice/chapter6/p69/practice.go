package main

// 闭包打印feibonacci数列
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		println(f())
	}
}

func fibonacci() func() int {
	var lastest, secondLastest int

	return func() int {
		if lastest == 0 && secondLastest == 0 {
			lastest, secondLastest = 1, 1
			return lastest
		}
		secondLastest, lastest = lastest, secondLastest+lastest
		return lastest
	}
}
