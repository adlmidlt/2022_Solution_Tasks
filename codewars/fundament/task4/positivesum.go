package main

func PositiveSum(numbers []int) int {

	sum := 0

	for _, num := range numbers {
		if num < 0 {
			num = 0
		} else {
			sum += num
		}
	}

	return sum
}
