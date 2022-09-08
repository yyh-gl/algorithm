package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := []int{1, 2, 2, 1, 2, 1}

	list := make([]*ListNode, len(input))
	for i := len(input); i > 0; i-- {
		var next *ListNode
		if i != len(input) {
			next = list[i]
		}

		list[i-1] = &ListNode{
			Val:  input[i-1],
			Next: next,
		}
	}

	fmt.Println(isPalindrome(list[0]))
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0, 50000)

	for head != nil {
		arr = append(arr, head.Val)

		head = head.Next
	}

	if len(arr)%2 == 1 {
		arr = append(arr[0:len(arr)/2], arr[len(arr)/2+1:]...)
	}

	reversed := reverse(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != reversed[i] {
			return false
		}
	}
	return true
}

func reverse(arr []int) []int {
	lenArr := len(arr)
	arr2 := make([]int, lenArr)
	for i := 0; i < lenArr/2; i++ {
		arr2[i], arr2[lenArr-i-1] = arr[lenArr-i-1], arr[i]
	}
	return arr2
}
