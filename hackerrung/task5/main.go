package main

import "fmt"

func main() {
	numbers := []int32{4, 9, 95, 93, 57, 4, 57, 93, 9}

	var numberDistinct []int32
	//var uniqueNum int32 = 0

	temp := make(map[int32]bool)

	// исключил повторные числа
	for _, numb := range numbers {
		if _, ok := temp[numb]; !ok {
			temp[numb] = true
			numberDistinct = append(numberDistinct, numb)
		}
	}
	var count int32
	for _, i := range numberDistinct {
		count = 0

		for _, j := range numbers {
			if i == j {
				count++
			}
		}
		fmt.Println(count)

		/*if count == 1 {
			uniqueNum = i
			break
		}*/
	}

	//fmt.Println(uniqueNum)
}
