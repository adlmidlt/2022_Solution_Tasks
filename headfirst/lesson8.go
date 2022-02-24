package main

import "fmt"

type subscribe struct {
	name   string
	rate   float64
	active bool
}

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

type student struct {
	name  string
	grade float64
}

func printInfoStudent(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func outStudent() {
	s := student{
		name:  "Alonzo Cole",
		grade: 92.3,
	}
	printInfoStudent(s)
}

func applyDiscount(s *subscribe) {
	s.rate = 4.99
}

func main() {
	fmt.Println("Hello lesson8")
	//subscriber1 := defaultSubscriber("Aman Sight")
	//applyDiscount(subscriber1)
	//printInfo(subscriber1)
	//
	//subscriber2 := defaultSubscriber("Beth Ryan")
	//printInfo(subscriber2)
	mainCar()

}

func printInfo(s *subscribe) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Rate: ", s.rate)
	fmt.Println("Active?: ", s.active)
}

func defaultSubscriber(name string) *subscribe {
	s := subscribe{
		name:   name,
		rate:   5.99,
		active: true,
	}
	return &s
}

func showInfo(p part) {
	fmt.Println(p.description)
	fmt.Println(p.count)
}

func getShowInfo(desc string) part {
	var p part
	p.description = desc
	return p
}

func ex1() {
	porsche := car{
		name:     "Porsche 911 R",
		topSpeed: 323,
	}
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	bolts := part{
		description: "Hex bolts",
		count:       24,
	}
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}

func perEx() {
	var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)
}

func subscribesEx() {
	subscribe1 := subscribe{
		name:   "Aman Singh",
		rate:   0,
		active: false,
	}
	fmt.Println(subscribe1)

	subscribe2 := subscribe{
		name:   "Aman Sight",
		rate:   0,
		active: false,
	}
	fmt.Println(subscribe2)

}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func mainCar() {
	mustang := car{
		name:     "Mustang Cobra",
		topSpeed: 255,
	}
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}
