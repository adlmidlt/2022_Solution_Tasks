package main

import (
	"fmt"
	"strconv"
)

func main() {
	var d int
	fmt.Scan(&d)
	str := strconv.Itoa(d)

	firstThreeNum := str[:3]
	lastSecondNum := str[3:]

	sumFirstThreeNum := 0
	sumLastThreeNum := 0

	for i, _ := range firstThreeNum {
		sumFirstThreeNum += int(firstThreeNum[i])
	}

	for i, _ := range lastSecondNum {
		sumLastThreeNum += int(lastSecondNum[i])
	}

	if sumFirstThreeNum == sumLastThreeNum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
