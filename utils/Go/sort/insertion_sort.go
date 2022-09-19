package sort

// 概要：指定要素をソート済みの部分の適切な場所に挿入していくことで並び替える
// 最良：O(N)
// 最悪：O(N^2)
// 平均：O(N^2)
// 与えられた配列以外の記憶領域が不要 (in-place)
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
