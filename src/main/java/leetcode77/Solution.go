package main

import "fmt"

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(combine(4, 2))
}
func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	if k <= 0 || n < k {
		return res
	}
	path := make([]int, 0)
	dfs(1, n, k, &path, &res)
	return res
}

func dfs(begin int, n int, k int, path *[]int, res *[][]int) {
	if len(*path) == k {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
		return
	}
	for i := begin; i <= n; i++ {
		*path = append(*path, i)
		dfs(i+1, n, k, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
