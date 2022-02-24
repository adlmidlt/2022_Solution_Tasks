package main

func SquareSum(numbers []int) int {
	// your code here
	result := 0
	for _, item := range numbers {
		result += item * item
	}
	return result
}
