package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cleanText := strings.TrimSpace(text)
	chars := []rune(cleanText)

	for i, _ := range chars {
		if chars[i] == chars[len(chars)-i-1] {
			fmt.Println("Палиндром")
			return
		} else {
			fmt.Println("Нет")
			return
		}
	}
}
