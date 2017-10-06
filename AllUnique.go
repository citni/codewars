package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(HasUniqueChar(" ab"))

}
func HasUniqueChar (str string) bool {

	split:= strings.Split(string(str),"")
	for i := 1; i < len(split); i++ {
		if (strings.Count(str, split[i]) > 1) {
			return false
			break
		}
	}
	return true
}