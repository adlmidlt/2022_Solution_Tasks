package main

import (
	"fmt"
)

// Define the Clock type here.

type Clock int

func New1(h, m int) Clock {
	minutesPerHour := h * 60
	totalAmountMinutes := (minutesPerHour + m) % 1440
	if totalAmountMinutes < 0 {
		totalAmountMinutes = 1440
	}

	return Clock(totalAmountMinutes)
}

func (c Clock) Add(m int) Clock {
	return New1(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New1(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func main() {
	clock2 := New1(24, 0)
	fmt.Println(clock2)

}
