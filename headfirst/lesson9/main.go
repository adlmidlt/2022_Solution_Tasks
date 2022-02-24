package main

import "fmt"

type Litters float64
type Milliliters float64
type Gallons float64

func main() {
	fmt.Println("Hello")

	carFuel := Gallons(10).ToLitters()
	busFuel := Litters(240).ToGallons()
	fmt.Println(carFuel)
	fmt.Println(busFuel)
	busFuel = Milliliters(20000).ToGallons()

	fmt.Println("G->L", carFuel, "litters", "\nL->G", busFuel, "gallons")

}

func (l Litters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) ToLitters() Litters {
	return Litters(g * 3.785)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

type Population int

func pool9() {
	population := Population(572)
	fmt.Println("Sleepy Creek Country population:", population)
	fmt.Println("Congratulation, Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek Country population:", population)
}

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

type Number int

func (n Number) add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func (n *Number) Double() {
	*n *= 2
}

func mainNumber() {
	ten := Number(10)
	ten.add(4)
	ten.subtract(5)
	four := Number(4)
	four.add(3)
	four.subtract(2)
	ten.Double()
	four.Double()
	fmt.Println(ten, four)
}
