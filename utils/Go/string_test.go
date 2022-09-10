package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_concatArr(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "1",
			args: []string{"a", "b", "c"},
			want: "abc",
		},
		{
			name: "2",
			args: []string{"a"},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := concatArr(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
