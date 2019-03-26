package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	// var i1 = 5
	// var i2 = 5
	// var intP1 *int = &i1
	// var intP2 *int = &i2
	// fmt.Printf("the i1 %d's pointer address is : %p\n", i1, intP1)
	// fmt.Printf("the i2 %d's pointer address is : %p\n", i2, intP2)
	// fmt.Printf("the pointer %p's value is %d", intP1, *intP1)

	// hex2byte(fmt.Sprintf("%p", intP1))
	// println(strings.Join([]string{"a", "b"}, ""))
	// println(7 / 3)
	hex2byte("0x17")
}

func hex2byte(hexStr string) string {
	//1.将字符串转换成10进制数

	//取出0x
	hexStr = strings.TrimLeft(hexStr, "0x")
	nums := strings.Split(hexStr, "")

	temp := 0

	for i := 0; i < len(nums); i++ {
		numStr := nums[len(nums)-1-i]
		var num int
		switch strings.ToLower(numStr) {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			num, _ = strconv.Atoi(numStr)
			break
		case "a", "b", "c", "d", "e", "f":
			num = int(strings.ToLower(numStr)[0]) - 97 + 10
		}
		temp += num * int(math.Pow(float64(16), float64(i)))
	}
	//2.将十进制数转换成二进制字符串
	result := ""
	println(temp)
	f(&temp, &result)

	println(result)

	return ""
}

func f(intValue *int, r *string) {
	if *intValue == 1 {
		*r = *r + "1"
		strSlice := strings.Split(*r, "")
		for from, to := 0, len(strSlice)-1; from < to; from, to = from+1, to-1 {
			strSlice[from], strSlice[to] = strSlice[to], strSlice[from]
		}
		*r = strings.Join(strSlice, "")
	} else {
		*r = *r + strconv.Itoa(*intValue%2)
		*intValue = *intValue / 2
		f(intValue, r)
	}
}
