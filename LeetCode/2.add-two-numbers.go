package main

/**
 * <p>You are given two <strong>non-empty</strong> linked lists representing two non-negative integers. The digits are stored in <strong>reverse order</strong>, and each of their nodes contains a single digit. Add the two numbers and return the sum&nbsp;as a linked list.</p>

<p>You may assume the two numbers do not contain any leading zero, except the number 0 itself.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg" style="width: 483px; height: 342px;" />
<pre>
<strong>Input:</strong> l1 = [2,4,3], l2 = [5,6,4]
<strong>Output:</strong> [7,0,8]
<strong>Explanation:</strong> 342 + 465 = 807.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> l1 = [0], l2 = [0]
<strong>Output:</strong> [0]
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
<strong>Output:</strong> [8,9,9,9,0,0,0,1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in each linked list is in the range <code>[1, 100]</code>.</li>
	<li><code>0 &lt;= Node.val &lt;= 9</code></li>
	<li>It is guaranteed that the list represents a number that does not have leading zeros.</li>
</ul>

**/
/**
 * [2,4,3]
[5,6,4]
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := make([]int, 0, 101)
	kuriage := false
	for l1 != nil || l2 != nil {
		switch {
		case l1 != nil && l2 != nil:
			tmp = append(tmp, l1.Val+l2.Val)
		case l1 != nil && l2 == nil:
			tmp = append(tmp, l1.Val)
		case l1 == nil && l2 != nil:
			tmp = append(tmp, l2.Val)
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if kuriage {
			tmp[len(tmp)-1] += 1
			kuriage = false
		}

		if tmp[len(tmp)-1] >= 10 {
			//if tmp[len(tmp)-1] == 10 {
			//	tmp = tmp[0 : len(tmp)-1]
			//}

			tmp[len(tmp)-1] -= 10
			kuriage = true
		}
	}

	if kuriage {
		tmp = append(tmp, 1)
	}

	var node, next *ListNode
	for i := len(tmp) - 1; i >= 0; i-- {
		node = &ListNode{
			Val:  tmp[i],
			Next: next,
		}
		next = node
	}
	return node
}
