package main

import (
	"errors"
	"fmt"
)

type WeightFodder interface {
	FodderAmount() (float64, error)
}

var ErrScaleMalfunction = errors.New("sensor error")

type SillyNephewError struct {
	AmountCows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.AmountCows)
}

func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderForCow, err := weightFodder.FodderAmount()
	fodderForOneCow := 0.0

	if err == ErrScaleMalfunction && fodderForCow > 0 {
		fodderForCow *= 2
	}

	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}

	if fodderForCow < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0, &SillyNephewError{AmountCows: cows}
	}

	fodderForOneCow = fodderForCow / float64(cows)

	return fodderForOneCow, nil
}
