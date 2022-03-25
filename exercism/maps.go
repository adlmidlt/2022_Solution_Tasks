package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if value, exist := units[unit]; !exist {
		return false
	} else if _, ok := bill[item]; !ok {
		bill[item] = value
	} else {
		bill[item] += value
	}

	return true
}

/*
Вернуть false, если данный элемент отсутствует в счете
Вернуть false, если данный юнит отсутствует на карте юнитов.
Верните false, если новое количество будет меньше 0.
Если новое количество равно 0, полностью удалите элемент из счета, а затем верните true.
В противном случае уменьшите количество товара и верните true.*/

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	vItem, existItem := bill[item]
	vUnit, existUnit := units[unit]
	if !existItem || !existUnit {
		return false
	}

	newQuantity := vItem - vUnit
	if newQuantity < 0 {
		return false
	} else {
		bill[item] -= units[unit]
	}

	if newQuantity == 0 {
		delete(bill, item)
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	vItem, existItem := bill[item]
	if !existItem {
		return 0, false
	}
	return vItem, true
}

func main() {
	bill := NewBill()
	units := Units()
	ok := RemoveItem(bill, units, "potato", "dozen")
	fmt.Println(ok, bill)
	// Output: false (because there are no carrots in the bill)
}
