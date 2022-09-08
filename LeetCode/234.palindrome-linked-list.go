package main

/**
 * <p>Given the <code>head</code> of a singly linked list, return <code>true</code><em> if it is a palindrome or </em><code>false</code><em> otherwise</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/03/03/pal1linked-list.jpg" style="width: 422px; height: 62px;" />
<pre>
<strong>Input:</strong> head = [1,2,2,1]
<strong>Output:</strong> true
</pre>

<p><strong>Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/03/03/pal2linked-list.jpg" style="width: 182px; height: 62px;" />
<pre>
<strong>Input:</strong> head = [1,2]
<strong>Output:</strong> false
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the list is in the range <code>[1, 10<sup>5</sup>]</code>.</li>
	<li><code>0 &lt;= Node.val &lt;= 9</code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you do it in <code>O(n)</code> time and <code>O(1)</code> space?
**/
/**
 * [1,2,2,1]
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
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

	reversed := reverse234(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != reversed[i] {
			return false
		}
	}
	return true
}

func reverse234(arr []int) []int {
	lenArr := len(arr)
	arr2 := make([]int, lenArr)
	for i := 0; i < lenArr/2; i++ {
		arr2[i], arr2[lenArr-i-1] = arr[lenArr-i-1], arr[i]
	}
	return arr2
}
