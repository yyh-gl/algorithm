package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utils_Max(t *testing.T) {
	type want struct {
		idx int
		num int
	}

	tests := []struct {
		nums []int
		want want
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: want{
				idx: 8,
				num: 9,
			},
		},
		{
			nums: []int{100, 2, 3, 4, 5, 6, 7, 8, 9},
			want: want{
				idx: 0,
				num: 100,
			},
		},
		{
			nums: []int{2, 3, 4, 100, 6, 7, 8, 9},
			want: want{
				idx: 3,
				num: 100,
			},
		},
		{
			nums: []int{1, 1, 1},
			want: want{
				idx: 0,
				num: 1,
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			idx, num := Max(tt.nums)
			assert.Equal(t, tt.want.idx, idx)
			assert.Equal(t, tt.want.num, num)
		})
	}
}

func Test_utils_Min(t *testing.T) {
	type want struct {
		idx int
		num int
	}

	tests := []struct {
		nums []int
		want want
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: want{
				idx: 0,
				num: 1,
			},
		},
		{
			nums: []int{100, 2, 3, 4, 5, 6, 7, 8, 1},
			want: want{
				idx: 8,
				num: 1,
			},
		},
		{
			nums: []int{2, 3, 4, 0, 6, 7, 8, 9},
			want: want{
				idx: 3,
				num: 0,
			},
		},
		{
			nums: []int{1, 1, 1},
			want: want{
				idx: 0,
				num: 1,
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			idx, num := Min(tt.nums)
			assert.Equal(t, tt.want.idx, idx)
			assert.Equal(t, tt.want.num, num)
		})
	}
}
