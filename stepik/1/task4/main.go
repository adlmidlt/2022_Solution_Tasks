package main

import "fmt"

func main() {
	var d int

	for i := 1; i <= 10; i++ {
		d = i * i
		fmt.Println(d)
	}
}
