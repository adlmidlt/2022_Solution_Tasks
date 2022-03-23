package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(val uint) uint {
		strNum := strconv.Itoa(int(val))
		var str string

		for i, item := range strNum {
			num, err := strconv.Atoi(string(item))
			if err != nil {
				fmt.Println(err.Error())
			}
			if num%2 == 0 && num != 0 {
				str += string(strNum[i])
			}
		}
		if str == "" {
			return 100
		}
		res, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		if res == 0 {
			return uint(100)
		}

		return uint(res)
	}
	/*var str string
		num := strconv.FormatUint(uint64(val), 10)

		for i := 0; i < len(num); i++ {
			if num[i]%2 == 0 && num[i] != 0 {
				str += string(num[i])
			}
		}
		fmt.Println(str)
		if str == "" {
			return uint(100)
		}
		res, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		if res == 0 {
			return uint(100)
		}

		return uint(res)
	}*/
	fmt.Println(fn(456034))
	fmt.Println(fn(1234))
	fmt.Println(fn(301))
	fmt.Println(fn(210))
	fmt.Println(fn(200))
	fmt.Println(fn(000))
}
