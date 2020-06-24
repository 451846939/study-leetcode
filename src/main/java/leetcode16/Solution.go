package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(threeSumClosest([]int{-1, 0, 1, 1, 55}, 3))
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, best := len(nums), math.MaxInt32

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(best-target) {
				best = sum
			}
			if sum > target {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else {
				l++
				if l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
