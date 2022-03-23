package main

import (
	"fmt"
	"strings"
)

type Battery interface {
	String() string
}

type BatteryCharge struct {
	charge string
}

func (b *BatteryCharge) String() string {
	zeroCount := strings.Count(b.charge, "0")
	oneCount := strings.Count(b.charge, "1")

	return fmt.Sprintf("[%s%s]", strings.Repeat(" ", zeroCount), strings.Repeat("X", oneCount))
}
func main() {
	var str string
	fmt.Scan(&str)

	batteryForTest := &BatteryCharge{charge: str}
	fmt.Println(batteryForTest)

}
