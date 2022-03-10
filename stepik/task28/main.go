package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)

	if (a+b)%2 == 0 {
		fmt.Printf("%d", (a+b)/2)
	} else {
		fmt.Printf("%.1f", float64(a+b)/2)
	}
}
