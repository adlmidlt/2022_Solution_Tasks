package main

import "fmt"

/*
3,213% за отрицательный баланс (баланс становится более отрицательным).
0,5% при положительном балансе менее 1000 долларов.
1,621% при положительном балансе больше или равном 1000 долларов и меньше 5000 долларов.
2,475% при положительном балансе больше или равном 5000 долларов.
*/

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	percent := 0.0
	if balance < 0.0 {
		percent = 3.213
	} else if balance >= 0.0 && balance < 1000 {
		percent = 0.5
	} else if balance > 0.0 && balance == 1000 && balance < 5000 {
		percent = 1.621
	} else if balance > 0.0 || balance == 5000 {
		percent = 2.475
	}

	return float32(percent)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	percentForProvidedBalance := balance / 100 * float64(InterestRate(balance))

	return percentForProvidedBalance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	sumPerYear := balance + Interest(balance)

	return sumPerYear
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	count := 0
	for balance <= targetBalance {
		balance = AnnualBalanceUpdate(balance)
		count++
	}

	return count
}

func main() {
	fmt.Println(YearsBeforeDesiredBalance(200.75, 214.88))
}
