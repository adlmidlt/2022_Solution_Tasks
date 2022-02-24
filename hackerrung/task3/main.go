package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	ss := strings.Split(s, ":")
	amPM := ss[2][2:4]
	ss[2] = ss[2][:2]
	switch amPM {
	case "AM":
		if ss[0] == "12" {
			ss[0] = "00"
		}
	case "PM":
		if ss[0] != "12" {
			h, _ := strconv.Atoi(ss[0])
			ss[0] = strconv.Itoa(h + 12)
		}
	}
	return ss[0] + ":" + ss[1] + ":" + ss[2]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
