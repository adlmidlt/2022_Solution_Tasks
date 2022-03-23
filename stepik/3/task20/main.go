package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sumNumber := 0

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
		sumNumber += number
	}
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(strconv.Itoa(sumNumber))
	if err != nil {
		panic(err.Error())
	}
	writer.Flush()
}
