package main

import "fmt"

/*
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word2 + word1)
	}
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = j
			}
			if j == 0 {
				dp[i][j] = i
			}
		}
	}
	for i := 1; i <= len(word2); i++ {
		for j := 1; j <= len(word1); j++ {
			if word1[j-1] != word2[i-1] {
				dp[i-1][j-1]++
			}
			dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]))
		}
	}
	return dp[len(word2)][len(word1)]
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
