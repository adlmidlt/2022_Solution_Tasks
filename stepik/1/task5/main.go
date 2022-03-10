package main

import "fmt"

func main() {
	var a int
	var b int
	var d int
	fmt.Scan(&a, &b)

	if a < b {
		for i := a; i <= b; i++ {
			d += i
		}
	}
	fmt.Println(d)
}
