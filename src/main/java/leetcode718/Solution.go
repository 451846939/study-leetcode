package main

import "fmt"

/*

给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例 1:

输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
说明:

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100
https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
*/
func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
func findLength(A []int, B []int) int {
	res := 0
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	//dp[i][j]表示第一个字符串前i个字符和第二个字符串前j个字符组成的最长公共字符串的长度
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func findLength2(A []int, B []int) int {
	res := 0
	dp := make([]int, len(B)+1)
	for i := 1; i <= len(A); i++ {
		for j := len(B); j >= 1; j-- {
			if A[i-1] == B[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}
			res = max(res, dp[j])
		}
	}
	return res
}
