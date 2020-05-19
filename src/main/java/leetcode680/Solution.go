package main

import "fmt"

/*
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(validPalindrome2("abca", 1))
}
func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
func validPalindrome2(s string, k int) bool {
	dp := make([][]int, 2)
	n := len(s)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = j
			} else {
				dp[i][j] = 0x3f3f3f3f
			}
		}
	}
	for i := 1; i <= n; i++ {
		dp[i&1][0] = i
		from, to := max(1, i-k), min(n, i+k)
		for j := from; j <= to; j++ {
			if s[i-1] == s[n-j] {
				dp[i&1][j] = dp[(i-1)&1][j-1]
			} else {
				dp[i&1][j] = min(dp[(i-1)&1][j]+1, dp[i&1][j-1]+1) // 从i - 1处和j - 1处转移
			}
		}
	}
	return dp[n&1][n] <= 2*k
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
