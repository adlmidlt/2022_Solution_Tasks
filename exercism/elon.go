package main

import "fmt"

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

type Track struct {
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      40,
		distance:     0,
	}
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	dist := Track{
		distance: distance,
	}

	return dist
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func (c *Car) Drive() {
	batteryConsumption := c.battery - c.batteryDrain
	if batteryConsumption < 0 {
		batteryConsumption = c.battery
	} else {
		c.distance += c.speed
		c.battery = batteryConsumption
	}

	fmt.Println(c)
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(track int) bool {
	for i := c.battery; i >= 0; i -= c.batteryDrain {
		c.distance += c.speed
		if track <= c.distance && i > 0 {
			return true
		}
	}

	return false
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	trackDistance := 100

	fmt.Println(car.CanFinish(trackDistance))
	// Output: true
}
