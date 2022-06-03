package main

import "fmt"

func main() {
	//got := romanToInt("III")
	//got := romanToInt("LVIII")
	got := romanToInt("MCMXCIV")
	fmt.Println(got)
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func romanToInt(s string) int {
	sum := 0
	num := 0
	prev := 0

	s = reverse(s)

	for _, v := range s {
		switch string(v) {
		case "M":
			num = 1000
		case "D":
			num = 500
		case "C":
			num = 100
		case "L":
			num = 50
		case "X":
			num = 10
		case "V":
			num = 5
		case "I":
			num = 1
		}

		if num < prev {
			sum -= num
		} else {
			sum += num
		}

		prev = num
	}
	return sum
}
