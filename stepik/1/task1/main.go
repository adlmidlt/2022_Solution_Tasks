package main

import (
	"fmt"
	"strconv"
)

func main() {
	var d int
	fmt.Scan(&d)

	str := strconv.Itoa(d)
	one := str[0]
	two := str[1]
	three := str[2]
	if one == two || one == three || two == three {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
