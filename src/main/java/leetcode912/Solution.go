package main

import "fmt"

/*
给你一个整数数组 nums，请你将该数组升序排列。



示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]


提示：

1 <= nums.length <= 50000
-50000 <= nums[i] <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(nums, p, r)
	quickSort(nums, p, q-1)
	quickSort(nums, q+1, r)
}

func partition(nums []int, p int, r int) int {
	pivot := nums[r]
	i := p
	for j := p; j < r; j++ {
		if nums[j] < pivot {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, r)
	return i
}
func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
