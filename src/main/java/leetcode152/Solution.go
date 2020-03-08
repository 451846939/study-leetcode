package main

import "fmt"

/*
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(maxProduct([]int{-2, 3, -4}))
	fmt.Println(maxProduct([]int{-2, -3, -4}))
	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2, -3}))
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct([]int{2, 3, 2, 4, 5}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4, 5}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	maxnum := nums[0]
	minnum := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := maxnum
		maxnum = max(minnum*nums[i], max(nums[i], maxnum*nums[i]))
		minnum = min(minnum*nums[i], min(nums[i], tmp*nums[i]))
		res = max(maxnum, res)
	}
	return res
}
func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}
func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
