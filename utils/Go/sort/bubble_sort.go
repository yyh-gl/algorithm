package sort

// 最良：O(N)
// 最悪：O(N^2)
// 平均：O(N^2)
// 与えられた配列以外の記憶領域が不要 (in-place)
func bubbleSort(nums []int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
