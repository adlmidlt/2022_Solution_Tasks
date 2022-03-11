package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)

	str := strconv.Itoa(a)

	for i := 0; i < len(str); i++ {
		number, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(number * number)
	}
}
