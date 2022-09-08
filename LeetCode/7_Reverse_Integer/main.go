package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(-9463847412))
}

func reverse(x int) int {
	isNegative := x < 0

	if isNegative {
		x = -1 * x
	}
	reversed := reverseInt(x)
	if isNegative {
		reversed = -1 * reversed
	}
	if -2147483648 <= reversed && reversed <= 2147483647 {
		return reversed
	}
	return 0
}
