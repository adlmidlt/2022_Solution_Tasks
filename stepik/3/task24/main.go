package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

/*type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}*/

//var s myStruct

func main() {
	m := user{
		Name:  "John Conor",
		Email: "test email",
	}
	data, _ := json.Marshal(m)
	data = bytes.Trim(data, "{")
	if !json.Valid(data) {
		fmt.Println("invalid json")
	}
	fmt.Printf("%s", data)
	/*data := []byte(`{"Name": "John Connor", "Age": 35, "Status": true, "Values": [15, 11, 37]}`)
	if err := json.Unmarshal(data, &s); err != nil {
		panic(err)
	}*/
	/*
		s := myStruct{
			Name:   "John Connor",
			Age:    35,
			Status: true,
			Values: []int{15, 11, 37},
		}

		data, err := json.MarshalIndent(s, "", "\t")
		if err != nil {
			panic(err)
		}*/
	//fmt.Printf("%v", data)
}
