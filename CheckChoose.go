package main

import (
	"fmt"
	"math/big"
)

const LIM = 64
var facts [LIM]uint64

func main() {
	fmt.Println("Result: ",CheckChoose(47129212243960,50))
}

func CheckChoose(m, n int) int {

	int64n := int64(n)
	int64m := int64(m)
	bigm := big.NewInt(int64m)

	first := new(big.Int).MulRange(1, int64n)
	second := new(big.Int)
	third := new(big.Int)

	product := new(big.Int)
	result := new(big.Int)

	for i := int64(0); i <= int64n ; i++ {

		second.MulRange(1, i)
		third.MulRange(1, int64n-i)

		product.Mul(second,third)
		result.Div(first, product)

		if bigm.Cmp(result) == 0 {
			inti := int(i)
			return inti
			break
		}
	}
	return -1
}