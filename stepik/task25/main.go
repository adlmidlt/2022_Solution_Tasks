package main

import "fmt"

const (
	hoursInDay       = 3600
	minutesInOneHour = 60
)

func main() {
	var seconds int
	fmt.Scan(&seconds)

	hours := seconds / hoursInDay
	minutes := (seconds - hours*hoursInDay) / minutesInOneHour

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
