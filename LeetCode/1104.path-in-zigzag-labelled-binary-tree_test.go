package main

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	type args struct {
		label int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	name: "1",
		//	args: args{label: 14},
		//	want: []int{1, 3, 4, 14},
		//},
		//{
		//	name: "2",
		//	args: args{label: 26},
		//	want: []int{1, 2, 6, 10, 26},
		//},
		{
			name: "3",
			args: args{label: 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		if got := pathInZigZagTree(tt.args.label); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. pathInZigZagTree() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
