package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "3",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "3",
			args: args{
				s:       "AB",
				numRows: 1,
			},
			want: "AB",
		},
	}
	for _, tt := range tests {
		if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
			t.Errorf("%q. convert() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
