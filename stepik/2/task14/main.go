package main

import (
	"fmt"
	"strings"
)

func main() {
	var textX string
	var textS string
	fmt.Scan(&textX, &textS)

	substr := strings.Contains(textX, textS)
	if substr {
		indexBegin := strings.Index(textX, textS)
		fmt.Println(indexBegin)
	} else {
		fmt.Println(-1)
	}
}
