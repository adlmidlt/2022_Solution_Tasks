package main

import "fmt"

func main() {
	var friendOfDima []string
	friendOfSemyon := make([]string, 3)
	friendOfDima = append(friendOfDima, "Костя", "Семён", "Таня")
	friendOfSemyon = append(friendOfSemyon, "Валера", "Таня", "Дима")
	friends := map[string][]string{
		"Dima":   friendOfDima,
		"Semyon": friendOfSemyon,
		"Костя":  nil,
	}
	_, value := friends["Костя"]
	fmt.Print(value, " ")
	delete(friends, "Dima")
	fmt.Print(friendOfSemyon[3])
}
