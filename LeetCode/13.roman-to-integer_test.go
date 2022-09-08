package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		if got := romanToInt(tt.args.s); got != tt.want {
			t.Errorf("%q. romanToInt() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
