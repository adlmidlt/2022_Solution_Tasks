package main

import (
	"fmt"
	"reflect"
)

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {

	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {

	for k, v := range r.Address {
		if k == "" && v == "" {
			return false
		}
		if k != "" && v == "" {
			return false
		}
	}

	if r.Name == "" || len(r.Address) == 0 {
		return false
	}

	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, resident := range residents {
		if resident.Name != "" && len(resident.Address) != 0 {
			count++
		}
	}
	return count
}

func main() {
	name := "Matthew Sanabria"
	age := 0
	address := map[string]string{}
	fmt.Println(reflect.ValueOf(address))

	resident := NewResident(name, age, address)

	fmt.Println(resident.HasRequiredInfo())
}
