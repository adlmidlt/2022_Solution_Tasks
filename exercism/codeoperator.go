package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var phoneNumber int
	_, err := fmt.Scan(&phoneNumber)
	if err != nil {
		fmt.Printf("Error should be number: %s", err.Error())
		os.Exit(1)
	}

	if math.Pow(10, 10) <= float64(phoneNumber) &&
		math.Pow(10, 11) > float64(phoneNumber) {
		str := strconv.Itoa(phoneNumber)
		codeOperator := str[1:4]
		number, err := strconv.Atoi(codeOperator)
		if err != nil {
			fmt.Printf("error of converting: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println(number)
	}
}
