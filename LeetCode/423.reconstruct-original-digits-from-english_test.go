package main

import "testing"

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: "owoztneoer"},
			want: "012",
		},
		{
			name: "2",
			args: args{s: "fviefuro"},
			want: "45",
		},
		{
			name: "3",
			args: args{s: "zerozero"},
			want: "00",
		},
		{
			name: "4",
			args: args{s: "oneone"},
			want: "11",
		},
		{
			name: "5",
			args: args{s: "oneonetwo"},
			want: "112",
		},
	}
	for _, tt := range tests {
		if got := originalDigits(tt.args.s); got != tt.want {
			t.Errorf("%q. originalDigits() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
