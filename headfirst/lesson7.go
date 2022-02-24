package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello lesson7")
	//learnMap()
	//example1()
	//nilWithMap()
	//okMap()
	//status("Anatolii")
	//status("Elena")
	//countLetter()
	//deleteMap()
	//hardProcessMap()
	pool3()
}

func pool3() {
	ranks := map[string]int{"gold": 1, "bronze": 3, "silver": 2}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank %d\n", medal, rank)
	}
}

func deleteMap() {
	var ok bool

	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v", prime, ok)
}

func countLetter() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}

func status(name string) {
	grades := map[string]float64{"Elena": 0, "Anatolii": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func okMap() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
}

func nilWithMap() {
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])
}

func example1() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace:", jewelry["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}

func learnMap() {
	ranks := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	elements := map[string]string{"H": "Hydrogen", "Li": "Lithium"}
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])

	isPrime := map[int]bool{1: true, 2: false}
	fmt.Println(isPrime[2])
	fmt.Println(isPrime[1])
}

func hardProcessMap() {
	lines, err := getString("votes.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}

func getString(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err = file.Close(); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
		return nil, scanner.Err()
	}

	return lines, nil
}
