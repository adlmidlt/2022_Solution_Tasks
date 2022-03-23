package main

import (
	"os"
	"strconv"
)

func main() {
	for i := 0; i <= 3; i++ {
		file, err := os.Create(strconv.Itoa(i) + "text.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

	}
	os.Rename("2text.txt", "4text.txt")
	for i := 4; i >= 2; i-- {
		os.Remove(strconv.Itoa(i) + "text.txt")
	}
}
