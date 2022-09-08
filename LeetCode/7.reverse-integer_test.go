package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{x: 12345},
			want: 54321,
		},
	}
	for _, tt := range tests {
		if got := reverse(tt.args.x); got != tt.want {
			t.Errorf("%q. reverse() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
