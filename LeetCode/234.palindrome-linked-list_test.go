package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.args.head); got != tt.want {
			t.Errorf("%q. isPalindrome() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
