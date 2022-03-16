package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dataFor := []byte("123")
	ioutil.WriteFile("test.txt", dataFor, 0600)
	ioutil.WriteFile("test.txt", dataFor, 0600)
	data, err := ioutil.ReadFile("text2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
