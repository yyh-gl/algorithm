package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode2
		l2 *ListNode2
	}
	tests := []struct {
		name string
		args args
		want *ListNode2
	}{
		{
			name: "success",
			args: args{
				l1: &ListNode2{
					Val: 2,
					Next: &ListNode2{
						Val: 4,
						Next: &ListNode2{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode2{
					Val: 5,
					Next: &ListNode2{
						Val: 6,
						Next: &ListNode2{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode2{
				Val: 7,
				Next: &ListNode2{
					Val: 0,
					Next: &ListNode2{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "success",
			args: args{
				l1: &ListNode2{
					Val: 9,
					Next: &ListNode2{
						Val: 9,
						Next: &ListNode2{
							Val: 9,
							Next: &ListNode2{
								Val: 9,
								Next: &ListNode2{
									Val: 9,
									Next: &ListNode2{
										Val: 9,
										Next: &ListNode2{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode2{
					Val: 9,
					Next: &ListNode2{
						Val: 9,
						Next: &ListNode2{
							Val: 9,
							Next: &ListNode2{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode2{
				Val: 8,
				Next: &ListNode2{
					Val: 9,
					Next: &ListNode2{
						Val: 9,
						Next: &ListNode2{
							Val: 9,
							Next: &ListNode2{
								Val: 0,
								Next: &ListNode2{
									Val: 0,
									Next: &ListNode2{
										Val: 0,
										Next: &ListNode2{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. addTwoNumbers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
