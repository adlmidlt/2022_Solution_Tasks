package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	if d%400 == 0 || (d%4 == 0 && d%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
