package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utils_reverse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := reverse("abc")
		assert.Equal(t, "cba", got)
	})
}
