package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var deleteN int
	fmt.Scan(&n, &deleteN)

	str := strconv.Itoa(n)

	for i := 0; i < len(str); i++ {
		number, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Println(err.Error())
		}
		if deleteN != number {
			fmt.Print(number)
		}
	}
}
