package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortString(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "1",
			args: []string{"c", "a", "b"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "2",
			args: []string{"c"},
			want: []string{"c"},
		},
		{
			name: "3",
			args: []string{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortString(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_sortInt(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "1",
			args: []int{3, 1, 2},
			want: []int{1, 2, 3},
		},
		{
			name: "2",
			args: []int{3},
			want: []int{3},
		},
		{
			name: "3",
			args: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortInt(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
