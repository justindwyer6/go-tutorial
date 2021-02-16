package main

import (
	"fmt"
	"strconv"
)

func main() {
	oneToTen := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, integer := range oneToTen {
		var evenOrOdd string
		if integer%2 == 0 {
			evenOrOdd = "even"
		} else {
			evenOrOdd = "odd"
		}
		fmt.Println(strconv.Itoa(integer) + " is " + evenOrOdd)
	}
}
