package main

func DNAtoRNA(dna string) string {
	chars := []rune(dna)
	var char string

	for i := 0; i < len(dna); i++ {
		if chars[i] == 'T' {
			chars[i] = 'U'
		}
		char = string(chars)
	}

	return char
}
