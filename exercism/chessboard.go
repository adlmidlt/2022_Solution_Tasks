package main

import "fmt"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	amount := 0
	for _, v := range cb[rank] {
		if v {
			amount++
		}
	}

	return amount
}

func newChessboard() Chessboard {
	return Chessboard{
		"A": Rank{true, false, true, false, false, false, false, true},
		"B": Rank{false, false, false, false, true, false, false, false},
		"C": Rank{false, false, true, false, false, false, false, false},
		"D": Rank{false, false, false, false, false, false, false, false},
		"E": Rank{false, false, false, false, false, true, false, true},
		"F": Rank{false, false, false, false, false, false, false, false},
		"G": Rank{false, false, false, true, false, false, false, false},
		"H": Rank{true, true, true, true, true, true, false, true},
	}
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	amount := 0

	if file > 0 && file < 9 {
		for _, v := range cb {
			if v[file-1] {
				amount++
			}
		}
	} else {
		return 0
	}

	return amount
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	amount := 0
	for _, v := range cb {
		for range v {
			amount++
		}
	}

	return amount
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	amount := 0
	for rank := range cb {
		amount += CountInRank(cb, rank)
	}

	return amount
}

func main() {
	fmt.Println(CountInRank(newChessboard(), "A"))
	fmt.Println(CountInFile(newChessboard(), 2))
}
