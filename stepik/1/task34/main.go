package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	for i := 1; i <= a; i *= 2 {
		fmt.Printf("%d ", i)
	}
}
