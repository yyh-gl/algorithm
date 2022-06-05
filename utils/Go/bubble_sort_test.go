package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utils_bubbleSort(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := bubbleSort([]int{3, 1, 4, 5, 2})
		assert.Equal(t, []int{1, 2, 3, 4, 5}, got)
	})
}
