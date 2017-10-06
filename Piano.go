package main

import (
	"fmt"
)

func main() {

	for i:= 1; i <= 10000; i+=1{
		fmt.Println(i, "NOTE:", BlackOrWhiteKey(i))
		if i % 88 == 0 {
			fmt.Println("---------------")
		}
	}
}

func BlackOrWhiteKey(keyPressCount int) string {

	n := []string{"white", "black", "white", "white", "black", "white", "black", "white", "white", "black", "white", "black"}

	i:= 0

	i = (keyPressCount - 1 ) % 88 % 12


	return n[i]
}
