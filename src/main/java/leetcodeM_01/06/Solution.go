package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*

字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：

字符串长度在[0, 50000]范围内。
https://leetcode-cn.com/problems/compress-string-lcci/
*/
func main() {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abbccd"))
}
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	var sb bytes.Buffer
	count := 1
	tmp := S[0]
	for i := range S {
		if i == 0 {
			continue
		}
		if S[i] == tmp {
			count++
		} else {
			sb.WriteByte(tmp)
			sb.WriteString(strconv.Itoa(count))
			count = 1
			tmp = S[i]
		}
		if i == len(S)-1 {
			sb.WriteByte(tmp)
			sb.WriteString(strconv.Itoa(count))
			count = 1
			tmp = S[i]
		}
	}
	res := sb.String()
	if len(res) > len(S) {
		return S
	}
	return res
}
