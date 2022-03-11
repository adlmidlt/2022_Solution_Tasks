package main

import "math"

var k, p, v float64

func T() float64 {
	w := W()
	t := 6 / w
	return t
}

func W() float64 {
	m := M()
	w := math.Sqrt(k / m)
	return w
}

func M() float64 {
	m := p * v
	return m
}
