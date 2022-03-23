package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	src := "aBcDeFg"
	test := "AbCdEfG"

	// Обратите внимание, что скобки после имени функции используются только при ее вызове
	src = strings.Map(invert, src)

	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)
}

func invert(r rune) rune {
	// Если буква строчная, то она возвращается заглавной
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	// Иначе возвращается строчной
	return unicode.ToLower(r)
}
