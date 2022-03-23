package main

import "fmt"

func main() {
	arr := [10]int{}
	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if value, inMap := m[arr[i]]; !inMap {
			m[arr[i]] = work(arr[i])
			fmt.Print(m[arr[i]], " ")
		} else {
			fmt.Print(value, " ")
		}
	}
}
