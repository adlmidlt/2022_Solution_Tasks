package main

import (
	"fmt"
	"time"
)

func main() {
	var strData string
	if _, err := fmt.Scan(&strData); err != nil {
		panic(err.Error())
	}
	data, err := time.Parse(time.RFC3339, strData)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(data.Format(time.UnixDate))
}
