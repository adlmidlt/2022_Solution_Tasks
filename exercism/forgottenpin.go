package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var pin int
	_, err := fmt.Scan(&pin)
	if err != nil {
		fmt.Printf("Error should be number: %s", err.Error())
		os.Exit(1)
	}
	str := strconv.Itoa(pin)

	if !strings.Contains(str, "0") {
		if str[0] != str[1] && str[0] != str[2] && str[1] != str[2] {
			number, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("error of converting: %s", err.Error())
				os.Exit(1)
			}

			one := number / 100
			two := number/10 - one*10
			three := number - number/10*10

			numbers := [][]int{
				{one, two, three},
				{one, three, two},
				{two, one, three},
				{two, three, one},
				{three, one, two},
				{three, two, one},
			}
			sorted(numbers)
			for _, num := range numbers {
				fmt.Println(num)
			}
		}
	}
}

func sorted(numbers [][]int) {
	sort.Slice(numbers[:], func(i, j int) bool {
		for x := range numbers[i] {
			if numbers[i][x] == numbers[j][x] {
				continue
			}
			return numbers[i][x] < numbers[j][x]
		}
		return false
	})
}
