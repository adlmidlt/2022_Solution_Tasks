package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type user struct {
	Students []struct {
		Rating []float64
	}
}

type AvgRating struct {
	Average float64
}

func main() {
	var u user

	data, _ := ioutil.ReadAll(os.Stdin)
	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println(err)
		return
	}

	if !json.Valid(data) {
		fmt.Println("invalid json")
	}

	avg := 0.0
	for i, _ := range u.Students {
		avg += float64(len(u.Students[i].Rating))
	}

	average := avg / float64(len(u.Students))
	avgRating := AvgRating{Average: average}

	dataOut, err := json.MarshalIndent(avgRating, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s", dataOut)
}
