package main

import (
	"fmt"
	"strings"
)

/*
Вам задана строка , состоящая из пробелов и латинских букв. Строка называется панграммой, если она содержит каждую из
26 латинских букв хотя бы раз. Определите является ли строка  панграммой.

Формат входных данных

В единственной строке входных данных записана строка  . Строка может содержать только пробелы, заглавные и строчные
буквы латинского алфавита. Заглавные и строчные буквы считаются одинаковыми.

Формат выходных данных

Выведите pangram если строка  является панграммой, иначе выведите not pangram.

Пример входных данных #1

We promptly judged antique ivory buckles for the next prize
Пример выходных данных #1

pangram
*/
func main() {
	str := "We promptly judged antique ivory buckles for the prize"

	if len(str) < 26 {
		fmt.Println("not pangram")
		return
	}

	lowerStr := strings.ToLower(str)
	alphabet := "abcdefghijklmnopqrstunwxyz"

	result := "pangram"

	for i := 0; i < 26; i++ {
		count := 0

		for a := 0; a < len(lowerStr); a++ {
			if alphabet[i] == lowerStr[a] {
				count++
			}
		}

		if count == 0 {
			result = "not pangram"
		}
	}

	fmt.Println(result)
}
