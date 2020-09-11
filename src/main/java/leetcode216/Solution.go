package main

/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	dfs(&ans, &tmp, k, 1, n)
	return ans
}

func dfs(ans *[][]int, tmp *[]int, k int, begin int, n int) {
	if n < 0 {
		return
	}
	if len(*tmp) == k && n == 0 {
		*ans = append(*ans, append([]int(nil), *tmp...))
		return
	}
	for i := begin; i <= 9; i++ {
		*tmp = append(*tmp, i)
		dfs(ans, tmp, k, i+1, n-i)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
