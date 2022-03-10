package main

import "fmt"

func main() {
	fmt.Println(fibonacci(4))
}

func fibonacci(n int) int {
	fPrev := 1
	fNext := 1

	for i := 0; i < n-2; i++ {
		fPrev, fNext = fNext, fPrev+fNext
	}

	return fNext
}
