package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	strArr := strings.Split(str, "")
	newText := strings.Join(strArr, "*")
	fmt.Println(newText)
}
