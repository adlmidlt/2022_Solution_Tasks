package main

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
}

func readTask() (interface{}, interface{}, interface{}) {
	var a interface{}
	var b interface{}
	var c interface{}
	switch a.(type), b.(type), c.(type) {
	case float64:
	default:

	}

	return nil, nil, nil
}
