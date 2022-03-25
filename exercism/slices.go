package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem1(slice []int, index int) int {
	for i := range slice {
		if i == index {
			return slice[i]
		}
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	var newArr []int
	for i := range slice {
		if i == index {
			slice[i] = value
		}
	}

	if index > len(slice)-1 || index < 0 {
		newArr = append(slice, value)
		return newArr
	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	var newArr []int
	newArr = append(newArr, value...)
	newArr = append(newArr, slice...)

	return newArr
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem1(slice []int, index int) []int {
	var newArr []int
	for i := range slice {
		if i == index {
			continue
		} else {
			newArr = append(newArr, slice[i])
		}
	}

	return newArr
}

func main() {
	//fmt.Println(FavoriteCards())
	//card := GetItem([]int{1, 2, 4, 1}, 2)
	//card := GetItem([]int{1, 2, 4, 1}, 10) // card == -1
	//fmt.Println(card)

	index := 4
	newCard := 1
	cards := SetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, index, newCard)
	fmt.Println(cards)

	//index := -1
	//newCard := 6
	//cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
	//fmt.Println(cards)

	//slice := []int{3, 2, 6, 4, 8}
	//cards := PrependItems(slice, 5, 1)
	//fmt.Println(cards)

	//slice := []int{3, 2, 6, 4, 8}
	//cards := PrependItems(slice)
	//fmt.Println(cards)

	//cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	//fmt.Println(cards)
	//cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
	//fmt.Println(cards)
}
