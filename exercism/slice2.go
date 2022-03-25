package main

import "fmt"

// TODO: define the 'PreparationTime()' function

const (
	OneLayerForNoodles = 50
	OneLayerForSauce   = 0.2
)

func PreparationTime(layers []string, avgTimeCookingInMinutes int) int {
	countOfLayers := len(layers)

	if avgTimeCookingInMinutes == 0 {
		avgTimeCookingInMinutes = 2
	}

	commonTimeCooking := countOfLayers * avgTimeCookingInMinutes

	return commonTimeCooking
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	countForNoodles := 0
	countForSauce := 0.0

	for _, item := range layers {
		if item == "noodles" {
			countForNoodles += OneLayerForNoodles
		}
		if item == "sauce" {
			countForSauce += OneLayerForSauce
		}
	}

	return countForNoodles, countForSauce
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(listOfIngredientFromFriend, myListIngredient []string) {

	for i := range myListIngredient {
		if myListIngredient[i] == "?" {
			myListIngredient[i] = listOfIngredientFromFriend[len(listOfIngredientFromFriend)-1]
		}
	}
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, countOfPortion int) []float64 {
	var amountOfIngredients []float64
	for _, item := range quantities {
		calcOfIngredientOnAmountPortion := float64(countOfPortion) / 2 * item
		amountOfIngredients = append(amountOfIngredients, calcOfIngredientOnAmountPortion)
	}

	return amountOfIngredients
}

func main() {
	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)
	// => []float64{ 2.4, 7.2, 21 }
}
