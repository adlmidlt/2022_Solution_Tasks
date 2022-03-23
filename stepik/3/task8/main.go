package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}

func adding(val1, val2 string) int64 {

	rune1 := []rune(val1)
	rune2 := []rune(val2)

	var str1 string
	var str2 string

	for _, item := range rune1 {
		if item <= 57 && item >= 48 {
			r1 := item
			str1 += string(r1)
		}
	}

	for _, item := range rune2 {
		if item <= 57 && item >= 48 {
			r2 := item
			str2 += string(r2)
		}
	}

	num1, err := strconv.Atoi(str1)
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := num1 + num2

	return int64(result)
}
