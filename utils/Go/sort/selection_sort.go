package sort

// 概要：配列から最小値を探し、配列の未ソート部分の先頭要素と入れ替えていくことで並べ替える
// 最良：O(N^2)
// 最悪：O(N^2)
// 平均：O(N^2)
// 与えられた配列以外の記憶領域が不要 (in-place)
func selectionSort(nums []int) []int {
	for i, _ := range nums {
		idx, _ := Min(nums[i:len(nums)])
		nums[i], nums[i+idx] = nums[i+idx], nums[i]
	}
	return nums
}

func Min(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}

	return
}
