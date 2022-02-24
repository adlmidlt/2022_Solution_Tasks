package main

import "fmt"

func main() {
	subscriber := Subscriber{
		Name:   "Aman Singh",
		Rate:   0,
		Active: false,
		HomeAddress: Address{
			Street:     "123 Oak St",
			City:       "Omaha",
			State:      "NE",
			PostalCode: "68111",
		},
	}
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("State:", subscriber.HomeAddress.State)
	fmt.Println("City:", subscriber.HomeAddress.City)
	fmt.Println("Street:", subscriber.HomeAddress.Street)
	fmt.Println("PostalCode:", subscriber.HomeAddress.PostalCode)

	employee := Employee{
		Name:   "Joy Carr",
		Salary: 60000,
		HomeAddress: Address{
			Street:     "456 Elm St",
			City:       "Portland",
			State:      "OR",
			PostalCode: "97222",
		},
	}
	fmt.Println("Name:", employee.Name)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("PostalCode:", employee.HomeAddress.PostalCode)
}
