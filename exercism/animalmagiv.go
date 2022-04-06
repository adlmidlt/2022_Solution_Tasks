package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {

	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a randomNumber int d with 1 <= d <= 20
func RollADie() int {
	SeedWithTime()

	num := rand.Intn(21)

	if num < 1 || num > 20 {
		num = 1
	}

	return num
}

// GenerateWandEnergy returns a randomNumber float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	num := rand.Intn(12)

	return float64(num)
}

// ShuffleAnimals returns a slice with all eight animal strings in randomNumber order
func ShuffleAnimals() []string {
	arrAnimal := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(arrAnimal), func(i, j int) {
		arrAnimal[i], arrAnimal[j] = arrAnimal[j], arrAnimal[i]
	})

	return arrAnimal
}

func main() {
	d := RollADie()
	fmt.Println(d)
}
