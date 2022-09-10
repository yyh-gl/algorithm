package main

/**
 * <p>The string <code>&quot;PAYPALISHIRING&quot;</code> is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)</p>

<pre>
P   A   H   N
A P L S I I G
Y   I   R
</pre>

<p>And then read line by line: <code>&quot;PAHNAPLSIIGYIR&quot;</code></p>

<p>Write the code that will take a string and make this conversion given a number of rows:</p>

<pre>
string convert(string s, int numRows);
</pre>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 3
<strong>Output:</strong> &quot;PAHNAPLSIIGYIR&quot;
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 4
<strong>Output:</strong> &quot;PINALSIGYAHRPI&quot;
<strong>Explanation:</strong>
P     I    N
A   L S  I G
Y A   H R
P     I
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;A&quot;, numRows = 1
<strong>Output:</strong> &quot;A&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consists of English letters (lower-case and upper-case), <code>&#39;,&#39;</code> and <code>&#39;.&#39;</code>.</li>
	<li><code>1 &lt;= numRows &lt;= 1000</code></li>
</ul>

**/
/**
 * "PAYPALISHIRING"
3
**/
func convert(s string, numRows int) string {
	sLen := len(s)
	if sLen == 1 || numRows == 1 {
		return s
	}

	rows := make([][]string, numRows)
	i := 0
	isAddMode := true
	for _, v := range s {
		rows[i] = append(rows[i], string(v))

		if i == numRows-1 {
			isAddMode = false
		} else if i == 0 {
			isAddMode = true
		}

		if isAddMode {
			i++
		} else {
			i--
		}
	}

	var res string
	for _, r := range rows {
		res += concatArr6(r)
	}
	return res
}

func concatArr6(arr []string) string {
	var s string
	for _, v := range arr {
		s += v
	}
	return s
}
