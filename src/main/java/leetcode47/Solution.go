package main

import "sort"

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	sort.Ints(nums)
	use := make([]bool, len(nums))
	dfs(nums, 0, len(nums), &use, &tmp, &res)
	return res
}

func dfs(nums []int, start int, size int, use *[]bool, tmp *[]int, ans *[][]int) {
	if start == size {
		t := make([]int, len(*tmp))
		copy(t, *tmp)
		*ans = append(*ans, t)
		return
	}
	for i := 0; i < size; i++ {
		if (*use)[i] || (i > 0 && nums[i] == nums[i-1] && !(*use)[i-1]) {
			continue
		}
		*tmp = append(*tmp, nums[i])
		(*use)[i] = true
		dfs(nums, start+1, size, use, tmp, ans)
		(*use)[i] = false
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
