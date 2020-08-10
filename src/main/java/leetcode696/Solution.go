package main

import "fmt"

/*
给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。

示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
示例 2 :

输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
注意：

s.length 在1到50,000之间。
s 只包含“0”或“1”字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-binary-substrings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {
	fmt.Println(countBinarySubstrings("00110"))
	fmt.Println(countBinarySubstrings("10101"))
}
func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	c1, c2, res := 1, 0, 0
	tmp := s[0]
	for i := 1; i < len(s); i++ {
		if tmp == s[i] {
			c1++
		} else {
			res += min(c1, c2)
			c2 = c1
			c1 = 1
			tmp = s[i]
		}
	}
	res += min(c1, c2)
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
