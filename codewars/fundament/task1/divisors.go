package main

import "fmt"

/*Count the number of divisors of a positive integer n.*/

func Divisors(n int) int {
	//Put your code here
	var count = 1

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}
	fmt.Println(count)
	return count
}
