package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	fPrev := 0
	fNext := 1
	n := 1

	if a == 0 {
		fmt.Println(0)
	} else if a == 1 {
		fmt.Println(1)
	} else {
		for fNext < a {
			fPrev, fNext = fNext, fPrev+fNext
			n++
		}
		if fNext == a {
			fmt.Println(n)
			return
		} else {
			fmt.Println(-1)
		}
	}
}
