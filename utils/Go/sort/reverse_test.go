package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utils_reverseString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "1",
			args: "abc",
			want: "cba",
		},
		{
			name: "2",
			args: "a",
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_utils_reverseIntArr(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "1",
			args: []int{26, 10, 6, 2, 1},
			want: []int{1, 2, 6, 10, 26},
		},
		{
			name: "2",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "3",
			args: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseIntArr(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_utils_reverseInt(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "1",
			args: 123378045,
			want: 540873321,
		},
		{
			name: "2",
			args: 1,
			want: 1,
		},
		{
			name: "3",
			args: 123,
			want: 321,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseInt(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
