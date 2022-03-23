package main

import "fmt"

func main() {
	a := 101.0 / 2.0
	fmt.Print(a + float64(uint8(a)))
}
