package main

const BatteryFull = 100

/*type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}*/

// NewCar creates a new remote controlled car with full battery and given specifications.
/*func NewCar(speed, batteryDrain int) Car {
	car := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      BatteryFull,
		distance:     0,
	}

	return car
}
*/
// NewTrack created a new track
/*func NewTrack(distance int) Track {
	dist := Track{
		distance: distance,
	}

	return dist
}*/

/*
Реализуйте функцию Drive, которая обновляет количество пройденных метров в зависимости от скорости автомобиля и
уменьшает заряд батареи в соответствии с разрядкой батареи:*/

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
/*func Drive(car Car) Car {
/*	dist := 0
	batteryConsumption := 0
	for i := BatteryFull; i <= car.battery; i -= car.batteryDrain {
		if i >= 1 {
			dist += car.speed
		} else {
			break
		}
		batteryConsumption = i
	}
	car.speed = dist*/
/*return Car{
		battery:      car.battery,
		speed:        car.speed,
		distance:     car.distance,
		batteryDrain: car.batteryDrain,
	}
}*/

/**
Чтобы закончить гонку, машина должна быть в состоянии проехать дистанцию гонки.
Это означает, что он не должен разряжать аккумулятор до пересечения финишной черты.
Реализуйте функцию CanFinish, которая принимает экземпляры Car и Track в качестве параметров и возвращает
true, если автомобиль может закончить гонку; в противном случае вернуть false.
Предположим, что машина только начинает гонку:
*/

// CanFinish checks if a car is able to finish a certain track.
/*func CanFinish(car Car, track Track) bool {

	for i := 0; i <= track.distance; i += car.speed {
		for j := BatteryFull; j >= car.battery; j -= car.batteryDrain {
			if i == track.distance && j != 0 {
				return true
			}
		}
	}

	return false
}*/

/*func main() {

speed := 5
batteryDrain := 5
car := NewCar(speed, batteryDrain)
car = Drive(car)
fmt.Println(car)*/
/*
	dist := 24
	raceTrack := NewTrack(dist)

	fmt.Println(CanFinish(car, raceTrack))

	fmt.Println(raceTrack)*/
//}
