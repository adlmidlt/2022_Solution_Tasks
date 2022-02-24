package main

func Litres(time float64) int {
	const drinkWater = 0.5

	result := drinkWater * time

	return int(result)
}
