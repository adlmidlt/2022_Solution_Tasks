package main

import (
	"fmt"
)

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	v1, ok := value1.(float64)
	if !ok {
		fmt.Printf("value=%v: %T\n", value1, value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		fmt.Printf("value=%v: %T\n", value2, value2)
		return
	}
	var result float64
	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			result = v1 + v2
			fmt.Printf("%.4f\n", result)
			return
		case "-":
			result = v1 - v2
			fmt.Printf("%.4f\n", result)
			return
		case "*":
			result = v1 * v2
			fmt.Printf("%.4f\n", result)
			return
		case "/":
			result = v1 / v2
			fmt.Printf("%.4f\n", result)
			return
		default:
			fmt.Print("неизвестная операция")
			return
		}
	} else {
		fmt.Println("неизвестная операция")
		return
	}
}

func readTask() (value1, value2, operation interface{}) {

	return 1.1, 1.2, "/"
}
