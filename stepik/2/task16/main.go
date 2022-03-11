package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	chars := []rune(str)

	for i := 0; i < len(chars); i++ {
		countStr := strings.Count(str, string(chars[i]))
		if countStr == 1 {
			fmt.Print(string(chars[i]))
		}
	}
}
