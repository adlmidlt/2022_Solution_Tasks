package main

import (
	"fmt"
	"unicode"
)

func main() {
	var passwd string
	fmt.Scan(&passwd)

	chars := []rune(passwd)
	str := ""

	if len(chars) >= 5 {
		for _, item := range chars {
			if !(unicode.Is(unicode.Latin, item) || unicode.Is(unicode.Digit, item)) {
				str = "Wrong password"
				break
			} else {
				str = "Ok"
			}
		}
	} else {
		str = "Wrong password"
	}
	fmt.Println(str)
}
