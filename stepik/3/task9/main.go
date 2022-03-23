package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}

	strNumbs := strings.Split(strings.TrimSpace(reader), ";")
	for i, item := range strNumbs {
		item = strings.ReplaceAll(item, " ", "")
		item = strings.ReplaceAll(item, ",", ".")
		strNumbs[i] = item
	}

	numb1, err := strconv.ParseFloat(strNumbs[0], 64)
	numb2, err := strconv.ParseFloat(strNumbs[1], 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := numb1 / numb2

	fmt.Printf("%.4f", result)
}
