package main

import (
	"errors"
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("should be from 1 to 64")
	}
	amountOfGrains := 0.0
	for i := 0; i < number; i++ {
		amountOfGrains = math.Pow(2, float64(i))
	}

	return uint64(amountOfGrains), nil
}

func Total() uint64 {
	return uint64(math.MaxUint64)
}

func main() {
	fmt.Println(Total())
}
