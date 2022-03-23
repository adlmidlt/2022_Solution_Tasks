package main

import "fmt"

func main() {
	m := map[int]int{
		1: 10,
	}
	if value, ok := m[1]; !ok {
		fmt.Println("Adsa")
	} else {
		fmt.Println(value)
	}

	var users []string
	fmt.Println(users)

}
