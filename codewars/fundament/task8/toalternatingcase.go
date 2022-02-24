package main

import (
	"unicode"
)

func ToAlternatingCase(str string) string {
	chars := []rune(str)
	var char []rune
	for _, item := range chars {
		if unicode.IsLower(item) {
			upperChar := unicode.ToUpper(item)
			char = append(char, upperChar)
		} else {
			lowerChar := unicode.ToLower(item)
			char = append(char, lowerChar)
		}
	}

	return string(char)
}
