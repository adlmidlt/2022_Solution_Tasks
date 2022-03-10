package main

import "fmt"

func main() {

	fmt.Println(vote(0, 0, 1))
}

func vote(x int, y int, z int) int {
	//print your code here
	arr := []int{x, y, z}

	result := 0
	countZero := 0
	countOne := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			countZero++
		} else {
			countOne++
		}
	}

	if countZero > countOne {
		result = 0
	} else {
		result = 1
	}

	return result
}
