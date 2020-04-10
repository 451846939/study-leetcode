package main

import (
	"bytes"
	"fmt"
)

/*
给定一个字符串，逐个翻转字符串中的每个单词。



示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


进阶：

请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(reverseWords(""))
}
func reverseWords(s string) string {
	l := 0
	r := 0
	q := make([]string, 0, len(s))
	for r < len(s) {
		if ' ' == s[r] {
			if r-l > 0 {
				q = append(q, s[l:r])
			}
			r++
			l = r
			continue
		}
		r++
	}
	if r != l {
		q = append(q, s[l:r])
	}
	var buffer bytes.Buffer
	for i := len(q) - 1; i >= 0; i-- {
		buffer.WriteString(q[i])
		if i != 0 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
