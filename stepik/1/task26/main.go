package main

import "fmt"

func main() {
	var a int
	var b int
	var c int

	fmt.Scan(&a, &b, &c)

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
