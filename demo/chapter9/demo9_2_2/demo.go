package main

import (
	"fmt"
	"math/big"
)

func main() {
	im := big.NewInt(9999)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)

	rm := big.NewRat(9999, 1111)
	ro := big.NewRat(22, 22)
	r := big.NewRat(1, 1)
	r.Mul(rm, ro).Add(r, ro).Sub(r, ro)
	fmt.Printf("Big Rat: %v\n", r)
}
