package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const formatData = "02.01.2006 15:04:05"

func main() {
	dataStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err.Error())
	}
	dataStr = strings.TrimSpace(dataStr)

	str := strings.Split(dataStr, ",")
	d1, err := time.Parse(formatData, str[0])
	d2, err := time.Parse(formatData, str[1])
	if err != nil {
		panic(err.Error())
	}
	dt1 := time.Since(d1)
	dt2 := time.Since(d2)

	if dt1 > dt2 {
		fmt.Println(dt1.Round(time.Second) - dt2.Round(time.Second))
	} else {
		fmt.Println(dt2.Round(time.Second) - dt1.Round(time.Second))
	}
}
