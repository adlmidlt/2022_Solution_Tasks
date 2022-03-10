package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := make([]int, a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])

		if i%2 == 0 {
			fmt.Print(b[i], " ")
		}
	}
}
