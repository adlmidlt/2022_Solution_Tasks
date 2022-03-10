package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	b := make([]int, a)
	var count = 0
	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])

		if b[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}
