package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	i := 0

	for i = b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}

	}
	fmt.Println("NO")
}
