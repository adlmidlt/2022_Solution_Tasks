package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello lesson10")

	event := Event{}
	if err := event.Date.SetYear(2022); err != nil {
		log.Fatal(err.Error())
	}

	if err := event.Date.SetMonth(2); err != nil {
		log.Fatal(err.Error())
	}

	if err := event.Date.SetDay(11); err != nil {
		log.Fatalf(err.Error())
	}

	if err := event.SetTitle("Mom's birthday"); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
