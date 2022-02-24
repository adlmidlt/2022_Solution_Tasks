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
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four
of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated
long integers.

Example: arr[1, 2, 3, 4, 5]

answer1 = 1 + 2 + 3 + 4
answer2 = 2 + 3 + 4 + 5

Output: 10 14
*/

func minMaxSum(arr []int32) {
	var sum int64
	minElem := arr[0]
	maxElem := arr[0]
	for _, item := range arr {
		sum += int64(item)
		if minElem > item {
			minElem = item
		}
		if maxElem < item {
			maxElem = item
		}
	}

	fmt.Printf("%d %d", sum-int64(maxElem), sum-int64(minElem))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	minMaxSum(arr)
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
