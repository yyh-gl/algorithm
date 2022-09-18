package utils

func Max(nums []int) (idx, n int) {
	n = nums[0]
	idx = 0
	for i, v := range nums {
		if n < v {
			n = v
			idx = i
		}
	}
	return
}

func Min(nums []int) (idx, n int) {
	n = nums[0]
	idx = 0
	for i, v := range nums {
		if n > v {
			n = v
			idx = i
		}
	}
	return
}
