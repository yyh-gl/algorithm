package utils

import "strconv"

// "hoge" => "egoh"
func reverseString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// [1,2,3] => [3,2,1]
func reverseIntArr(arr []int) []int {
	lenArr := len(arr)
	arr2 := make([]int, lenArr)
	for i := 0; i < lenArr/2; i++ {
		arr2[i], arr2[lenArr-i-1] = arr[lenArr-i-1], arr[i]
	}
	return arr2
}

// 123 => 321
func reverseInt(n int) int {
	str := strconv.Itoa(n)
	reversedStr := ""
	for _, v := range str {
		reversedStr = string(v) + reversedStr
	}
	num, _ := strconv.Atoi(reversedStr)
	return num
}
