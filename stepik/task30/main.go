package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a)

	countMin := 0
	b := 0
	for i := 0; i < a; i++ {
		fmt.Scan(&arr[i])
		b = arr[i]
	}

	min := b

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if min == arr[i] {
			countMin++
		}
	}

	fmt.Println("countMin = ", countMin)
}
