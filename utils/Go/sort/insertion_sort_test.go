package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utils_insertionSort(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := insertionSort([]int{3, 1, 4, 5, 2})
		assert.Equal(t, []int{1, 2, 3, 4, 5}, got)
	})

	t.Run("success", func(t *testing.T) {
		got := insertionSort([]int{5, 4, 3, 2, 1})
		assert.Equal(t, []int{1, 2, 3, 4, 5}, got)
	})
}
