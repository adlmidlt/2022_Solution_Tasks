package main

import "strings"

func DuplicateEncode(word string) string {
	result := ""
	lowerWordCase := strings.ToLower(word)

	for i, _ := range lowerWordCase {
		if strings.LastIndex(lowerWordCase, string(lowerWordCase[i])) == strings.Index(lowerWordCase, string(lowerWordCase[i])) {
			result += "("
		} else {
			result += ")"
		}
	}

	return result
}

// 2 способ решения
func DuplicateEncode1(word string) (r string) {
	word = strings.ToLower(word)

	t := map[rune]int{}
	for _, c := range word {
		t[c]++
	}

	for _, c := range word {
		if t[c] == 1 {
			r += "("
		} else {
			r += ")"
		}
	}

	return
}
