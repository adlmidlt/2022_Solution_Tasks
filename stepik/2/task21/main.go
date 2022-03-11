package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)

	c := float64(a*a + b*b)
	fmt.Println(math.Sqrt(c))
}
