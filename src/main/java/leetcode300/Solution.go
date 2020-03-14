package main

import "fmt"

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/
func main() {
	//fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

/*动态规划时间复杂度为o(n^2)*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//定义状态dp[i]表示第i个数结尾最长上升子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1
	maxlen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			tmp := 1
			//如果dp[j]为0初始化为1
			dp[j] = max(tmp, dp[j])
			if nums[i] > nums[j] {
				//dp[j]的最长上升子序列长度+1
				tmp = dp[j] + 1
			}
			dp[i] = max(tmp, dp[i])
			//因为dp并不存在一个递增的关系，所以需要记录最大值，每次判断
			maxlen = max(dp[i], maxlen)
		}
	}
	return maxlen
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
