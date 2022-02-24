package main

import "strings"

func RemoveChar(word string) string {
	chars := []rune(word)
	result := ""

	for i := 0; i < len(word); i++ {
		if i == 0 {
			chars[i] = ' '
		}
		if i == len(word)-1 {
			chars[i] = ' '
		}
		result = strings.TrimSpace(string(chars))
	}
	return result
}

// RemoveChar1 - 2 способ
func RemoveChar1(word string) string {
	return word[1 : len(word)-1]
}
