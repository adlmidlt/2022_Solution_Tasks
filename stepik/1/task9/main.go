package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var res string

	for {
		fmt.Scan(&a)
		if a < 10 {
			continue
		} else if a > 100 {
			break
		} else {
			res += strconv.Itoa(a) + "\n"
		}
	}
	fmt.Println(res)
}
