package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const upperLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Robot struct {
	NameRobot string
}

// randomName Random names robots.
func randomName() (string, error) {
	name := fmt.Sprintf("%s%s%s",
		randomLetter(upperLetter),
		randomLetter(upperLetter),
		strconv.Itoa(randomNumber(100, 999)))

	return name, nil
}

func (r *Robot) Name() (string, error) {
	if r.NameRobot == "" {
		nameRobot, err := randomName()
		if err != nil {
			return "", err
		}
		r.NameRobot = nameRobot
	}

	return r.NameRobot, nil
}

// randomNumber Random number in range.
func randomNumber(from, to int) int {
	if from > to {
		return from
	}

	return rand.Intn(to-from) + from
}

// randomLetter Random choice letter and number.
func randomLetter(source string) string {
	return string(source[rand.Intn(len(source))])
}

func main() {
	randomName()
}
