package main

import (
	"fmt"
)

func main() {

	for i:= 1; i <= 10000; i+=1{
		fmt.Println(i, "NOTE:", WhichNote(i))
		if i % 88 == 0 {
		fmt.Println("---------------")
		}
	}
}

func WhichNote(keyPressCount int) string {

	n := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

	i:= 0

	i = (keyPressCount - 1 ) % 88 % 12


	return n[i]
}
