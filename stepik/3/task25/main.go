package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func main() {

	newUser := user{
		Name:     "Alex",
		Email:    "email@dsadsa.tu",
		Status:   true,
		Language: []byte("ru"),
	}
	data, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	newUser.Language = []byte("en")
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(newUser.Language))
}
