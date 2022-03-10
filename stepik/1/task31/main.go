package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)

	str := strconv.Itoa(a)
	sum := 0
	numSum := 0
	for _, item := range str {
		number, _ := strconv.Atoi(string(item))
		sum += number
	}

	strSum := strconv.Itoa(sum)
	for _, item := range strSum {
		numberSum, _ := strconv.Atoi(string(item))
		numSum += numberSum
	}
	fmt.Println(numSum)
}
