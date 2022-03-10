package main

import "fmt"

func main() {
	var a int
	var sum int

	fmt.Scan(&a)
	b := make([]int, a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])
		if b[i] > 9 && b[i] < 100 && b[i]%8 == 0 {
			sum += b[i]
		}
	}
	fmt.Println(sum)
}
