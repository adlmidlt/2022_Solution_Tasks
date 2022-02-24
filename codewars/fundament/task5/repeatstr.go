package main

import "strings"

/*
Write a function called repeatStr which repeats the given string 'string' exactly n times.
*/

func RepeatStr(repetitions int, value string) string {

	result := ""

	for i := 0; i < repetitions; i++ {
		result += value
	}

	return result
}

// RepeatStr1 - 2-й способ
func RepeatStr1(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}
