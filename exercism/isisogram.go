package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	noSpaceInWord := strings.TrimSpace(word)
	wordUpper := strings.ToUpper(noSpaceInWord)

	for i := range wordUpper {
		for j := len(wordUpper) - 1; j > i; j-- {
			if wordUpper[i] == wordUpper[j] && unicode.IsLetter(rune(wordUpper[i])) {
				return false
			}
		}
	}

	for i, r := range wordUpper {
		if unicode.IsLetter(r) && strings.ContainsRune(wordUpper[i+1:], r) {
			return false
		}
	}

	// Можно и так но я решил разобраться как это работает.
	/*for i, value := range wordUpper {
		if i < strings.LastIndex(wordUpper, string(value)) {
			return false
		}
	}*/
	return true
}

func main() {
	fmt.Println(IsIsogram("six-year-old"))
}
