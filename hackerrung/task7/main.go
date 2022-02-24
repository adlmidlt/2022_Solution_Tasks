package main

import (
	"fmt"
)

func main() {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	var sumFirstDiagonal int32 = 0
	var sumSecondDiagonal int32 = 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				sumFirstDiagonal += arr[i][j]
			}
			if (i + j) == (len(arr) - 1) {
				sumSecondDiagonal += arr[i][j]
			}
		}
	}

	var moduleExpression int32 = 0
	moduleExpression = sumFirstDiagonal - sumSecondDiagonal
	if moduleExpression < 0 {
		moduleExpression = 0 - moduleExpression
	}
	fmt.Println(moduleExpression)
}

/*func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	leftToRight := 0
	rightToLeft := 0

	for index := range arr {
		leftToRight += arr[index][index]
		rightToLeft += arr[index][len(arr)-1-index]
	}

	fmt.Printf("result: %d", func(x int) int {
		if x < 0 {
			return 0 - x
		}
		return x
	}(leftToRight-rightToLeft))
}*/
