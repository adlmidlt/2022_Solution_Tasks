package main

import "fmt"

func changer(i *int) {
	*i = 2 * *i
}
func main() {
	i := 11
	changer(&i)
	fmt.Println(i)
}
