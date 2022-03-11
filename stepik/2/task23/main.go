package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)

	chars := []rune(a)
	max := ""
	for i := range chars {
		if max < string(chars[i]) {
			max = string(chars[i])

		}
	}

	fmt.Println(max)
}
