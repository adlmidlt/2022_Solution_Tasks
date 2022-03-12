package main

import "fmt"

func printNumber(i int) {
	i = g(i)
	fmt.Println(i*3 - 10)
}

func null() {
	fmt.Println("123")
}
