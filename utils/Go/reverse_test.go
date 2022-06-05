package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utils_reverseString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := reverseString("abc")
		assert.Equal(t, "cba", got)
	})
}
