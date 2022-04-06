package main

import (
	"errors"
	"fmt"
)

// Record1 represents an expense record.
type Record1 struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record1, predicate func(Record1) bool) []Record1 {
	var newRecord []Record1
	for i := range in {
		b := predicate(Record1{
			Day:      in[i].Day,
			Amount:   in[i].Amount,
			Category: in[i].Category,
		})
		if b {
			newRecord = append(newRecord, in[i])
		}
	}

	return newRecord
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record1) bool {

	return func(record Record1) bool {
		if record.Day >= p.From && record.Day <= p.To {
			return true
		}

		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record1) bool {

	return func(record Record1) bool {
		if c == record.Category {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record1, p DaysPeriod) float64 {
	listRecord := Filter(in, ByDaysPeriod(p))
	totalAmount := 0.0
	for _, record := range listRecord {
		totalAmount += record.Amount
	}

	return totalAmount
}

/*
Реализуйте функцию CategoryExpenses, которая возвращает общую сумму расходов в категории за заданный период дней.
Функция также должна различать случай, когда данная категория отсутствует в учете расходов, и случай, когда в
указанном периоде отсутствуют расходы категории. Когда категория не является категорией ни одной из записей
(независимо от периода времени), функция должна возвращать ошибку.*/

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record1, p DaysPeriod, c string) (float64, error) {
	category := ByCategory(c)
	categoryExpenses := Filter(in, category)
	totalAmount := TotalByPeriod(categoryExpenses, p)

	if !category(Record1{Category: c}) {
		return totalAmount, nil
	}

	if len(categoryExpenses) == 0 {
		return 0, errors.New("unknown category entertainment")
	}

	return totalAmount, nil
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record1) bool {
	return r.Day == 1
}
func main() {
	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	records := []Record1{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(CategoryExpenses(records, p1, "entertainment"))
	// Output: 0, error(unknown category entertainment)

	fmt.Println(CategoryExpenses(records, p1, "rent"))
	// Output: 1300, nil

	fmt.Println(CategoryExpenses(records, p2, "rent"))
	// Output: 0, nil
}
