package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 0))
}

func sumInt(num ...int) (int, int) {
	sum := 0

	var arr1 []int

	arr1 = append(arr1, num...)

	for _, item := range arr1 {
		sum += item
	}

	return len(arr1), sum
}
