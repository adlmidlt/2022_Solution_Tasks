package main

import "fmt"

/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой
последовательности, которые равны ее наибольшему элементу.

Формат входных данных:

Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность
не входит, а служит как признак ее окончания).*/

func main() {
	var a int
	var max int
	var d int

	fmt.Scan(&a)

	for a != 0 {
		if a > max {
			max = a
			d = 1
		} else if a == max {
			d += 1
		}
		fmt.Scan(&a)
	}
	fmt.Println(d)
}
