package main

func Solution(word string) string {
	result := ""

	for _, item := range word {
		result = string(item) + result
	}

	return result
}
