package main

import (
	"strconv"
	"strings"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	decodeString("100[leetcode]")
}
func decodeString(s string) string {
	res := ""
	numStack := make([]int64, 0)
	stringStack := make([]string, 0)
	runes := []rune(s)
	num := ""
	for j := 0; j < len(runes); j++ {
		if runes[j] >= '0' && runes[j] <= '9' {
			num += string(runes[j])
		} else if runes[j] == '[' {
			n, _ := strconv.ParseInt(num, 10, 32)
			numStack = append(numStack, n)
			num = ""
			stringStack = append(stringStack, res)
			res = ""
		} else if runes[j] == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			ss := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			res = ss + strings.Repeat(res, int(num))
		} else {
			res += string(runes[j])
		}
	}
	return res
}
