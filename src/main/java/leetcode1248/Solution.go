package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 k。

如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。

请返回这个数组中「优美子数组」的数目。



示例 1：

输入：nums = [1,1,2,1,1], k = 3
输出：2
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
示例 2：

输入：nums = [2,4,6], k = 1
输出：0
解释：数列中不包含任何奇数，所以不存在优美子数组。
示例 3：

输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
输出：16


提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10^5
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	//输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
	//
	//输出：16
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
func numberOfSubarrays(nums []int, k int) int {
	/*
		我们定义 pre[i]pre[i] 为 [0..i][0..i] 中奇数的个数，则 pre[i] 可以由 pre[i−1] 递推而来，即：

		pre[i]=pre[i−1]+(nums[i]&1)

		那么「[j..i]这个子数组里的奇数个数恰好为 k 」这个条件我们可以转化为

		pre[i]-pre[j-1]==k

		简单移项可得符合条件的下标 j 需要满足

		pre[j−1]==pre[i]−k

	*/
	dp := make([]int, len(nums)+1)
	if len(nums) == 0 || len(nums) < k {
		return 0
	}
	res := 0
	odd := 0
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			odd++
		}
		if odd >= k {
			res += dp[odd-k]
		}
		dp[odd] += 1
	}
	return res
}
