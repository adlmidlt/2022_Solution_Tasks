package main

import (
	"fmt"
	"strings"
)

func Score(word string) int {
	letterValues := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}
	wordUpper := strings.ToUpper(word)
	amountOfPoint := getAmountOfPoint(wordUpper, letterValues)

	return amountOfPoint
}

// getAmountOfPoint Get amount of point.
func getAmountOfPoint(word string, letterValues map[int][]string) int {
	counter := 0
	for value, letterValue := range letterValues {
		for i := range letterValue {
			for j := range word {
				if letterValue[i] == string(word[j]) {
					counter += value
				}
			}
		}
	}

	return counter
}

func main() {
	fmt.Println(Score("recovered"))
}
