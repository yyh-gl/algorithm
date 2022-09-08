package main

import "strconv"

/**
 * <p>Given a signed 32-bit integer <code>x</code>, return <code>x</code><em> with its digits reversed</em>. If reversing <code>x</code> causes the value to go outside the signed 32-bit integer range <code>[-2<sup>31</sup>, 2<sup>31</sup> - 1]</code>, then return <code>0</code>.</p>

<p><strong>Assume the environment does not allow you to store 64-bit integers (signed or unsigned).</strong></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> x = 123
<strong>Output:</strong> 321
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> x = -123
<strong>Output:</strong> -321
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> x = 120
<strong>Output:</strong> 21
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= x &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

**/
/**
 * 123
**/
func reverse(x int) int {
	isNegative := x < 0

	if isNegative {
		x = -1 * x
	}
	reversed := reverseInt(x)
	if isNegative {
		reversed = -1 * reversed
	}
	if -2147483648 <= reversed && reversed <= 2147483647 {
		return reversed
	}
	return 0
}

func reverseInt(n int) int {
	str := strconv.Itoa(n)
	reversedStr := ""
	for _, v := range str {
		reversedStr = string(v) + reversedStr
	}
	num, _ := strconv.Atoi(reversedStr)
	return num
}
