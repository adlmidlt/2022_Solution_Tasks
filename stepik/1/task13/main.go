package main

import "fmt"

func main() {
	workArray := [10]uint8{}
	var a uint8
	var b1, b2 uint8

	for i := range workArray {
		fmt.Scan(&a)
		workArray[i] = a
	}

	for j := 0; j < 3; j++ {
		fmt.Scan(&b1, &b2)
		workArray[b1], workArray[b2] = workArray[b2], workArray[b1]
	}

	for _, item := range workArray {
		fmt.Printf("%d ", item)
	}
}
