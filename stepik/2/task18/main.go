package main

import (
	"errors"
	"fmt"
)

func main() {

	n, err := divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(n)
	}

}

func divide(a int, b int) (int, error) {
	return a / b, errors.New("ошибка")
}
