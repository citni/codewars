package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Integral is %f\n", Simpson(2))
}

func Simpson(n int) float64 {

	h := (math.Pi) / n

	result := (math.Pi/(3*n))(math.Sin()^3)

	return float64(n)
}