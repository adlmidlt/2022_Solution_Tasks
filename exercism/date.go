package main

import (
	"fmt"
	"time"
)

const (
	dateFormatOne   = "1/2/2006 15:04:05"
	dateFormatTwo   = "January 2, 2006 15:04:05"
	dateFormatThree = "Monday, January 2, 2006 15:04:05"
	dateFormatFour  = "Monday, January 2, 2006, at 15:04"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	convDate, err := time.Parse(dateFormatOne, date)
	if err != nil {
		panic(err.Error())
	}
	return convDate
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()

	convDate, err := time.Parse(dateFormatTwo, date)
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
	convDate, err := time.Parse(dateFormatThree, date)
	if err != nil {
		panic(err.Error())
	}
	if convDate.Hour() >= 12 && convDate.Hour() < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	convDate := Schedule(date)
	formatDate := convDate.Format(dateFormatFour)
	return fmt.Sprintf("You have an appointment on %s.", formatDate)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	date := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)

	return date
}

func main() {

}
