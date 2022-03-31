package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func Valid(id string) bool {
	sum := 0

	if len(id) == 0 || len(id) == 1 {
		return false
	}

	numbers := make([]int, 0, len(id))

	for _, r := range id {
		if unicode.IsNumber(r) {

			number, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err.Error())
			}
			numbers = append(numbers, number)
		} else {
			return false
		}

		for i, number := range numbers {
			result := 0
			if i%2 == 0 {
				result = number * 2
				if result > 9 {
					result -= 9
				}
				sum += result
			} else {
				sum += number
			}
		}
	}
	fmt.Println(numbers)
	return sum%10 == 0

}

func main() {
	fmt.Println(Valid("059"))
	fmt.Println(Valid("55 444 285"))
}
