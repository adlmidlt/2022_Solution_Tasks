package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func main() {
	const URL = "https://github.com/semyon-dev/stepik-go/raw/master/work_with_files/task_sep_1/task.data"

	rs, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	rd := bufio.NewReader(rs.Body)

	i := 1
	for {
		s, err := rd.ReadString(';')
		if err == io.EOF {
			break
		}

		if s == "0;" {
			fmt.Println(i)
			break
		}
		i++
	}
}
