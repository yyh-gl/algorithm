package utils

func concatArr(arr []string) string {
	var s string
	for _, v := range arr {
		s += v
	}
	return s
}
