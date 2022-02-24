package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
Print the decimal value of each fraction on a new line with  places after the decimal.

There are n=5 elements, two positive, two negative and one zero.
Their ratios are, 2/5=0.400000, 2/5=0.400000 and 1/5=0.200000. Results are printed as:

Example: arr = [1, 1, 0, -1, -1]

positive(1) = count(2)
zero(0) = count(1)
negative(-1) = count(2)

Output:

0.400000
0.400000
0.200000
*/
func plusMinus(arr []int32) {
	var positive = 0.0
	var negative = 0.0
	var zero = 0.0

	countElements := len(arr)
	for i := 0; i < countElements; i++ {
		if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative++
		} else if arr[i] == 0 {
			zero++
		}
	}
	fmt.Printf("%.6f\n", positive/float64(countElements))
	fmt.Printf("%.6f\n", negative/float64(countElements))
	fmt.Printf("%.6f\n", zero/float64(countElements))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32
	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
