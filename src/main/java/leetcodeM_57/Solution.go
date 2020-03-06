package main

/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	findContinuousSequence(9)
}
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	if target <= 2 {
		return res
	}
	left := 1
	right := 1
	sum := 0
	for right < target {
		if sum == target {
			tmp := make([]int, right-left)
			for i, j := 0, left; i < right-left; i, j = i+1, j+1 {
				tmp[i] = j
			}
			res = append(res, tmp)
			sum -= left
			left++
		} else if sum > target {
			sum -= left
			left++
		} else {
			sum += right
			right++
		}
	}
	return res
}
