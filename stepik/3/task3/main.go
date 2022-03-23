package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  {"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: {"Миллионик"},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Город":     100,
		"Миллионик": 1000,
		"Село":      100,
	}

	for key := range cityPopulation {
		for _, city := range groupCity[100] {
			if city == key {
				break
			}
			delete(cityPopulation, key)
		}
	}
	fmt.Println(cityPopulation)
}
