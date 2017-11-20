package main

import (
	"fmt"
)

func main() {
	fmt.Println(MaxBall(15))
}

func MaxBall(v0 int) int {
	const g = float32(9.81)
	v := float32(v0)/ float32(3.6)
	t := float32(float32(v) / g)
	return int(int64(t * 10 + 0.5))
}