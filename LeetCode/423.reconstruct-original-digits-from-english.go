package main

import (
	"strconv"
)

/**
 * <p>Given a string <code>s</code> containing an out-of-order English representation of digits <code>0-9</code>, return <em>the digits in <strong>ascending</strong> order</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> s = "owoztneoer"
<strong>Output:</strong> "012"
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> s = "fviefuro"
<strong>Output:</strong> "45"
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s[i]</code> is one of the characters <code>[&quot;e&quot;,&quot;g&quot;,&quot;f&quot;,&quot;i&quot;,&quot;h&quot;,&quot;o&quot;,&quot;n&quot;,&quot;s&quot;,&quot;r&quot;,&quot;u&quot;,&quot;t&quot;,&quot;w&quot;,&quot;v&quot;,&quot;x&quot;,&quot;z&quot;]</code>.</li>
	<li><code>s</code> is <strong>guaranteed</strong> to be valid.</li>
</ul>

**/
/**
 * "owoztneoer"
**/
func originalDigits(s string) string {
	counts := make([]int, 10)
	for _, v := range s {
		switch string(v) {
		case "z":
			counts[0]++
		case "w":
			counts[2]++
		case "x":
			counts[6]++
		case "g":
			counts[8]++
		case "u":
			counts[4]++
		case "s":
			counts[7]++
		case "f":
			counts[5]++
		case "h":
			counts[3]++
		case "i":
			counts[9]++
		case "o":
			counts[1]++
		}
	}

	counts[7] -= counts[6]
	counts[5] -= counts[4]
	counts[3] -= counts[8]
	counts[9] -= counts[5] + counts[6] + counts[8]
	counts[1] -= counts[0] + counts[2] + counts[4]

	var res string
	for i, c := range counts {
		if c > 0 {
			for j := 0; j < c; j++ {
				res += strconv.Itoa(i)
			}
		}
	}
	return res
}
