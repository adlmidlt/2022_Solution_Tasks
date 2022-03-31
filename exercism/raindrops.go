package main

import "strconv"

func Convert(number int) string {
	res := ""
	if number%3 == 0 {
		res += "Pling"
	}

	if number%5 == 0 {
		res += "Plang"
	}

	if number%7 == 0 {
		res += "Plong"
	}

	if number%7 != 0 && number%5 != 0 && number%3 != 0 {
		str := strconv.Itoa(number)
		return str
	}
	return res
}

func main() {

}
