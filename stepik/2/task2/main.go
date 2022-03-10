package main

import "fmt"

func main() {

	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	arr := [4]int{}

	min := 0

	for i := 0; i < 4; i++ {
		fmt.Scan(&arr[i])
		min = arr[i]
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}
