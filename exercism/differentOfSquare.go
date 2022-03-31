package main

import (
	"fmt"
	"math"
)

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	squareOfSum := math.Pow(float64(sum), 2)

	return int(squareOfSum)
}

func SumOfSquares(n int) int {
	sumOfSquare := 0.0
	for i := 1; i <= n; i++ {
		sumOfSquare += math.Pow(float64(i), 2)
	}

	return int(sumOfSquare)
}

func Difference(n int) int {
	different := SquareOfSum(n) - SumOfSquares(n)

	return different
}

func main() {
	fmt.Println(SquareOfSum(10))
	fmt.Println(SumOfSquares(10))
	fmt.Println(Difference(10))
}
