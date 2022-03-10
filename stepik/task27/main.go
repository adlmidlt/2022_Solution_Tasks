package main

import "fmt"

func main() {
	var a int
	var b int
	var c int

	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
