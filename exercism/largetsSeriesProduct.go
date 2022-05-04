package main

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if len(digits) < span {
		return 0, errors.New("span needs to be smaller the digits")
	}

	if span < 0 {
		return 0, errors.New("span needs to be positive")
	}

	var bigNumber int64
	for i := 0; i <= len(digits)-span; i++ {

		var num int64 = 1
		for j := 0; j < span; j++ {

			number, err := strconv.Atoi(string(digits[j+i]))
			if err != nil {
				return 0, err
			}

			num *= int64(number)
		}

		if bigNumber <= num {
			bigNumber = num
		}
	}
	return bigNumber, nil
}

func main() {
	LargestSeriesProduct("29", 2)
}
