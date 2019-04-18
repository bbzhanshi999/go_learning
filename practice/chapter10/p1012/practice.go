package main

import (
	"fmt"
)

// 给定结构体类型 T:

// ```go
// type T struct {
//     a int
//     b float32
//     c string
// }
// ```

// 值 `t`: `t := &T{7, -2.35, "abc\tdef"}`。给 T 定义 `String()`，使得 `fmt.Printf("%v\n", t)` 输出：`7 / -2.350000 / "abc\tdef"`。

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	// return strconv.FormatInt(int64(t.a), 10) + " / " + strconv.FormatFloat(float64(t.b), 'f', 2, 32) + " / " + t.c
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}
