package main

import (
	"fmt"
	"log"
	"time"
)

const LayoutFormat = "1/02/2006 15:04:05"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	convDate, err := time.Parse(LayoutFormat, date)
	if err != nil {
		panic(err.Error())
	}
	return convDate
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()

	convDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err.Error())
	}

	if now.Unix() > convDate.Unix() {
		return true
	}

	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}

func main() {
	vals := []string{"7/25/2019 13:45:00"}

	for _, val := range vals {
		t, err := time.Parse("1/02/2006 15:04:05", val)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(t.UTC())
	}
}
