package main

import "fmt"

func main() {
	var x int // вклад
	var p int // процент
	var y int // итого
	var a int // лет
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	for i := 1; x <= y; i++ {
		x += x * p / 100
		a = i
	}
	fmt.Println(a)
}
