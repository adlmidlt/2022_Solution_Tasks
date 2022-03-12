package main

import (
	"fmt"
	"headfirst/stepik/2/task7/stepik"
)

func main() {
	a := 100
	LetsGo(a)
	fmt.Print(a, " ")
	stepik.LetsGo(a)
}

func LetsGo(b int) {
	b++
}
