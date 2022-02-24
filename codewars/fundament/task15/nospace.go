package main

func NoSpace(word string) string {

	result := ""

	for _, item := range word {
		if item != ' ' {
			result += string(item)
		}
	}
	return result
}
