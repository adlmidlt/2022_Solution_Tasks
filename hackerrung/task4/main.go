package main

import "fmt"

/*
There is a collection of input strings and a collection of query strings. For each query string, determine how many
times it occurs in the list of input strings. Return an array of the results.
*/

func main() {
	query := []string{"aba", "xzxb", "ab"}
	input := []string{"aba", "baba", "aba", "xzxb"}

	var result []int32

	for _, queryItem := range query {
		var occurrences int32 = 0

		for _, inputItem := range input {
			if inputItem == queryItem {
				occurrences += 1
			}
		}

		result = append(result, occurrences)
	}

	fmt.Printf("%v", result)
}
