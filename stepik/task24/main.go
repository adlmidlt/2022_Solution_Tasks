package main

import (
	"fmt"
	"strconv"
)

func main() {
	var threeDigitNumber int
	fmt.Scan(&threeDigitNumber)

	str := strconv.Itoa(threeDigitNumber)

	i1, i2, i3 := str[2:], str[1:2], str[:1]

	fmt.Print(i1, i2, i3)
}
