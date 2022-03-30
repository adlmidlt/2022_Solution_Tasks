package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	distanceOfHamming := 0

	if len(a) != len(b) {
		return 0, errors.New("strings not equal")
	}

	for i := range b {
		if a[i] == b[i] || a == "" {
			continue
		}
		distanceOfHamming++
	}

	return distanceOfHamming, nil
}

func main() {
	fmt.Println(Distance("GGACTGAAATCTG", "GGACTGAAATCTG"))
}
