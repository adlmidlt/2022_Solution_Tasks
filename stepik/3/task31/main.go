package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	dataStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err.Error())
	}
	dataStr = strings.TrimSpace(dataStr)

	compile := regexp.MustCompile(`(\d+) мин\. (\d+) сек\.`)
	text := compile.FindStringSubmatch(dataStr)

	minutes, err := strconv.Atoi(text[1])
	seconds, err := strconv.Atoi(text[2])
	if err != nil {
		panic(err.Error())
	}

	data := time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)

	fmt.Println(time.Unix(now, 0).Add(data).UTC().Format(time.UnixDate))
}
