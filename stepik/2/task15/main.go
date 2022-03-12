package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)

	chars := []rune(str)

	for i := 0; i < len(chars); i++ {
		if i%2 != 0 {
			fmt.Print(string(chars[i]))
		}
	}
}
