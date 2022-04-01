package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	deleteSpace := strings.TrimSpace(id)

	if len(deleteSpace) <= 1 {
		return false
	}

	numbers := make([]int, 0, len(deleteSpace))
	for i := 0; i < len(deleteSpace); i++ {
		if string(deleteSpace[i]) == " " {
			continue
		}

		if unicode.IsNumber(rune(deleteSpace[i])) {
			number, _ := strconv.Atoi(string(deleteSpace[i]))
			numbers = append(numbers, number)
		} else {
			return false
		}
	}
	sum := getSumOfNumbers(numbers)

	return sum%10 == 0
}

// getSumOfNumbers Get sum of numbers.
func getSumOfNumbers(numbers []int) int {
	var sum int
	for j, number := range numbers {
		dl := len(numbers) - j
		if dl%2 == 0 {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
	}

	return sum
}

func main() {
	//fmt.Println(Valid("4539 3195 0343 6467")) // true
	//fmt.Println(Valid("1"))                   // false
	//fmt.Println(Valid("0"))                   // false
	//fmt.Println(Valid("059")) // true
	//fmt.Println(Valid("059a"))                // false
	//fmt.Println(Valid("055-444-285"))         // false
	//fmt.Println(Valid("0000 0"))              // true
	//fmt.Println(Valid("091"))                 // true
	//fmt.Println(Valid("109"))                 // true
	fmt.Println(Valid("095 245 88")) // true
}
