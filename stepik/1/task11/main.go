package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)

	str1 := strconv.Itoa(a)
	str2 := strconv.Itoa(b)

	for i, _ := range str1 {
		for j, _ := range str2 {
			if str1[i] == str2[j] {
				fmt.Print(string(str1[i]) + " ")
			}
		}
	}
}
