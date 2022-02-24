package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello lesson5")
	average()
	//pool()
	//readFile()
}

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if err = file.Close(); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if scanner.Err() != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return numbers, nil
}

func average() {
	numbers, err := getFloats("data.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	var sum float64

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)
	fmt.Println(os.Args)
}

func readFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func pool() {
	numbers := [6]int{3, 16, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number < 20 {
			fmt.Println(i, number)
		}
	}
}
