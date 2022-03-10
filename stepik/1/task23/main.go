package main

import (
	"fmt"
	"strconv"
)

func main() {
	var threeDigitNumber int
	fmt.Scan(&threeDigitNumber)

	str := strconv.Itoa(threeDigitNumber)

	var sum int

	for _, item := range str {
		number, err := strconv.Atoi(string(item))
		sum += number
		fmt.Println(err.Error())
	}

}
