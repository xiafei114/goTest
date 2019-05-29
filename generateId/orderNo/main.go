package main

import (
	"fmt"
	"gotest/generateId/orderNo/str"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(str.MakeYearDaysRand(10))
	}
}
