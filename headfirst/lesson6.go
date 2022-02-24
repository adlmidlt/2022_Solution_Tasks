package main

import "fmt"

func main() {
	fmt.Println("Hello lesson6")
	//prime()
	//pool1()
	//arrays()
	//slices()
	//sliceNext()
	//severalInts(1, 2, 3, 4, 5)
	//fmt.Println(inRange(25, 100, 21, 34, 54, 120))
	//fmt.Println(inRange(-25, 40, -21, 34, 54, 120))
	//pool2()
	//fmt.Println(average1(100, 50))
	//fmt.Println(average1(90.7, 89.7, 98.5, 92.3))
	example()
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func example() {
	intSlice := []int{1, 2, 3}
	severalInts(intSlice...)
	stringSlice := []string{"a", "b", "c", "d"}
	mix(1, true, stringSlice...)
}

func pool2() {
	fmt.Println(sum(7, 9))
	fmt.Println(sum(1, 4, 2))
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func average1(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func inRange(min, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func sliceNext() {
	var intSlice []int
	var stringSlice []string
	fmt.Printf("%#v\n", intSlice)
	fmt.Printf("%#v\n", stringSlice)

	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])
}

func slices() {
	slice := []string{"a", "b"}
	fmt.Println(slice)

	slice = append(slice, "c")
	fmt.Println(slice, len(slice))

	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	slice[0] = "XX"
	fmt.Println(slice)
}

func arrays() {
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)

	array2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := array2[2:5]
	array2[2] = "X"
	fmt.Println(array2)
	fmt.Println(slice2)

	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	array3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice3)
	fmt.Println(slice4)
}

func pool1() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string{"a", "b", "c"}
	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}

func prime1() {
	primes := make([]int, 5)
	notes := make([]string, 7)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(len(primes))
	fmt.Println(len(notes))
}
