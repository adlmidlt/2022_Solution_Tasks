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
	car := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     5,
	}

	return car
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
func Drive(car Car) Car {
	batteryConsumption := car.battery - car.batteryDrain
	if batteryConsumption < 0 {
		batteryConsumption = car.battery
	} else {
		car.distance += car.speed
	}

	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      batteryConsumption,
		distance:     car.distance,
	}
}

func CanFinish(car Car, track Track) bool {
	for i := car.battery; i >= 0; i -= car.batteryDrain {
		car.distance += car.speed
		if track.distance < car.distance && i > 0 {
			return true
		}
	}

	return false
}

func main() {
	speed := 2
	batteryDrain := 6
	car := NewCar(speed, batteryDrain)
	car = Drive(car)

	dist := 80
	raceTrack := NewTrack(dist)

	fmt.Println(CanFinish(car, raceTrack))
}
