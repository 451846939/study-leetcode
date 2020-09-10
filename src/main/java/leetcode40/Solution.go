package main

import "sort"

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func combinationSum2(candidates []int, target int) [][]int {
	tmp := make([]int, 0)
	res := make([][]int, 0)
	sort.Ints(candidates)
	dfs(candidates, 0, target, &res, &tmp)
	return res
}

func dfs(candidates []int, index int, target int, res *[][]int, tmp *[]int) {

	if target == 0 {
		t := make([]int, len(*tmp))
		copy(t, *tmp)
		*res = append(*res, t)
		return
	}
	for i := index; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		*tmp = append(*tmp, candidates[i])
		dfs(candidates, i+1, target-candidates[i], res, tmp)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
