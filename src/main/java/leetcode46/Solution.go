package main

import "fmt"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(nums, &res, 0)
	return res
}
func backtrack(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		n := make([]int, len(nums))
		copy(n, nums)
		*res = append(*res, n)
	}
	for i := index; i < len(nums); i++ {
		//交换元素
		nums[index], nums[i] = nums[i], nums[index]
		//递归
		backtrack(nums, res, index+1)
		//取消交换
		nums[index], nums[i] = nums[i], nums[index]
	}

}
