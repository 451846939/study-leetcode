package main

import "math"

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	dp := make([][]int, r+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[r-1][c] = 0
	dp[r][c-1] = dp[r-1][c]
	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[0][0]
}
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	dp := make([][]int, r)
	for i := range dp {
		dp[i] = make([]int, c)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < r; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < c; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[r-1][c-1]
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
