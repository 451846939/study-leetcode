package main

import "fmt"

/*

给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
https://leetcode-cn.com/problems/longest-palindrome/

*/
func main() {
	//	babaabab
	fmt.Println(longestPalindrome("abccccdd"))
}
func longestPalindrome(s string) int {
	//counts := make(map[uint8]int, len(s))
	counts := make([]int, 58)
	for i := range s {
		//counts[s[i]]++
		counts[s[i]-'A']++
	}
	count := 0
	for _, v := range counts {
		count += v >> 1 << 1
		if v%2 == 1 && count%2 == 0 {
			count++
		}
	}
	return count
}
func longestPalindrome2(s string) int {
	//counts := make(map[uint8]int, len(s))
	counts := make([]int, 58)
	for i := range s {
		//counts[s[i]]++
		counts[s[i]-'A']++
	}
	count := 0
	for _, v := range counts {
		count += v >> 1 << 1
	}
	//s如果是奇数count也是奇数，s是偶数count也是偶数就直接返回，否则说明有奇数加一即可
	if count == len(s) {
		return count
	}
	return count + 1
}
