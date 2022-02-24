package main

import "strings"

/*
Example: (input --> output)

"ATTGC" --> "TAACG"
"GTAT" --> "CATA"
Example:
dnaStrand []        `shouldBe` []
dnaStrand [A,T,G,C] `shouldBe` [T,A,C,G]
dnaStrand [G,T,A,T] `shouldBe` [C,A,T,A]
dnaStrand [A,A,A,A] `shouldBe` [T,T,T,T]
*/

func DNAStrand(dna string) string {
	chars := []rune(dna)
	var char string
	for i := 0; i < len(chars); i++ {
		if chars[i] == 'A' {
			chars[i] = 'T'
		} else if chars[i] == 'T' {
			chars[i] = 'A'
		}
		if chars[i] == 'G' {
			chars[i] = 'C'
		} else if chars[i] == 'C' {
			chars[i] = 'G'
		}
		char = string(chars)
	}
	return char
}

// 2 способ
var dnaReplacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

func DNAStrand1(dna string) string {
	return dnaReplacer.Replace(dna)
}

func main() {
	DNAStrand("AAGA")
}
