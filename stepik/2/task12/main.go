package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cleanText := strings.TrimSpace(text)
	chars := []rune(cleanText)

	if unicode.IsUpper(chars[0]) && chars[len(chars)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
