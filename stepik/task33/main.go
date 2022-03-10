package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	str := "korov"

	switch a <= 99 {
	case a == 1 || a == 21 || a == 31 || a == 41 || a == 51 || a == 61 || a == 71 || a == 81 || a == 91:
		fmt.Printf("%d %sa", a, str)
	case a <= 4 || (a >= 2 && a <= 2) || (a >= 22 && a <= 24) || (a >= 32 && a <= 34) || (a >= 42 && a <= 44) || (a >= 52 && a <= 54) ||
		(a >= 62 && a <= 64) || (a >= 72 && a <= 74) || (a >= 82 && a <= 84) || (a >= 92 && a <= 94):
		fmt.Printf("%d %sy", a, str)
	case (a >= 5 && a <= 20) || (a >= 25 && a <= 30) || (a >= 35 && a <= 40) || (a >= 45 && a <= 50) || (a >= 55 && a <= 60) ||
		(a >= 65 && a <= 70) || (a >= 75 && a <= 80) || (a >= 85 && a <= 90) || (a >= 95 && a <= 99):
		fmt.Printf("%d %s", a, str)
	}
}

/*var a int
fmt.Scan(&a)

switch {
case a%10 ==1 && a%100 != 11: fmt.Println(a,"korova")
case a%10 == 2 && a%100 != 12: fallthrough
case a%10 == 3 && a%100 != 13: fallthrough
case a%10 == 4 && a%100 != 14: fmt.Println(a,"korovy")
default:
fmt.Println(a,"korov")
}*/

/*
1 - корова
2-4 - коровы
5-20 - коров

21 - корова
22-24 - коровы
25-30 - коров

31 - корова
32-34 - коровы
35-40 - коров

41 - корова
42-44 - коровы
45-50 - коров*/
