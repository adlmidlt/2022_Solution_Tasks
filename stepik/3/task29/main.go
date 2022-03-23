package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const formatData = "2006-01-02 15:04:05"

func main() {
	strData, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	strData = strings.TrimSpace(strData)
	data, err := time.Parse(formatData, strData)
	if err != nil {
		panic(err.Error())
	}

	if data.Hour() >= 13 {
		data = data.Add(time.Hour * 24)
	}
	fmt.Println(data.Format(formatData))
}
